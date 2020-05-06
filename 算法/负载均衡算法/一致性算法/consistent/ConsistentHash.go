package consistent

import (
	"crypto/sha1"
	"math"
	"sort"
	"strconv"
	"sync"
)

const (
	//DefaultVirualSpots default virual spots
	DefaultVirualSpots = 400
)

type Node struct {
	NodeKey   string
	spotValue uint32
}

type NodesArray []Node

func (p NodesArray) Len() int           { return len(p) }
func (p NodesArray) Less(i, j int) bool { return p[i].spotValue < p[j].spotValue }
func (p NodesArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p NodesArray) Sort()              { sort.Sort(p) }

//HashRing store nodes and weights
type HashRing struct {
	virualSpots int
	Nodes       NodesArray
	weights     map[string]int
	mu          sync.RWMutex
}

// 初始化虚拟环
//NewHashRing create a hash ring with virual spots
func NewHashRing(spots int) *HashRing {
	if spots == 0 {
		spots = DefaultVirualSpots
	}

	h := &HashRing{
		virualSpots: spots,
		weights:     make(map[string]int),
	}
	return h
}

// 批量添加节点
//AddNodes add nodes to hash ring
func (h *HashRing) AddNodes(nodeWeight map[string]int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	// 计算每个服务器权重
	for nodeKey, w := range nodeWeight {
		h.weights[nodeKey] = w
	}
	h.generate()
}

// 单个添加节点
//AddNode add node to hash ring
func (h *HashRing) AddNode(nodeKey string, weight int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.weights[nodeKey] = weight
	h.generate()
}

//RemoveNode remove node
func (h *HashRing) RemoveNode(nodeKey string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.weights, nodeKey)
	h.generate()
}

//UpdateNode update node with weight
func (h *HashRing) UpdateNode(nodeKey string, weight int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.weights[nodeKey] = weight
	h.generate()
}

func (h *HashRing) generate() {
	var totalW int
	// 计算出总权重
	for _, w := range h.weights {
		totalW += w
	}

	totalVirtualSpots := h.virualSpots * len(h.weights)
	h.Nodes = NodesArray{}

	for nodeKey, w := range h.weights {
		spots := int(math.Floor(float64(w) / float64(totalW) * float64(totalVirtualSpots)))
		// 计算出hash值
		for i := 1; i <= spots; i++ {
			hash := sha1.New()
			hash.Write([]byte(nodeKey + ":" + strconv.Itoa(i)))
			hashBytes := hash.Sum(nil)
			n := Node{
				NodeKey:   nodeKey,
				spotValue: genValue(hashBytes[6:10]),
			}
			h.Nodes = append(h.Nodes, n)
			hash.Reset()
		}
	}
	h.Nodes.Sort()
}

func genValue(bs []byte) uint32 {
	if len(bs) < 4 {
		return 0
	}
	v := (uint32(bs[3]) << 24) | (uint32(bs[2]) << 16) | (uint32(bs[1]) << 8) | (uint32(bs[0]))
	return v
}

//GetNode get node with key
func (h *HashRing) GetNode(s string) string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if len(h.Nodes) == 0 {
		return ""
	}

	hash := sha1.New()
	hash.Write([]byte(s))
	hashBytes := hash.Sum(nil)
	v := genValue(hashBytes[6:10])
	i := sort.Search(len(h.Nodes), func(i int) bool { return h.Nodes[i].spotValue >= v })

	if i == len(h.Nodes) {
		i = 0
	}
	return h.Nodes[i].NodeKey
}
