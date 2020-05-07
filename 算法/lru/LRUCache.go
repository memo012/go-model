package main

import "fmt"

/*
	最近最未使用删除算法： LRU算法
*/
type LRUCache struct {
	CacheCapacity int               // 容量大小
	caches        map[int]*CacheNode // 缓存值
	first         *CacheNode        // 第一个缓存值
	last          *CacheNode        // 最后一个缓存值
}

type CacheNode struct {
	key   int
	value string
	pre   *CacheNode
	next  *CacheNode
}

// 初始化
func Constructor(size int) *LRUCache {
	return &LRUCache{
		CacheCapacity: size,
		caches:        make(map[int]*CacheNode, size),
		first:         nil,
		last:          nil,
	}
}

func (c *LRUCache) put(k int, v string) {
	elemValue, ok := c.caches[k]
	if ok {
		// moveToFirst()
		elemValue.value = v
		c.moveToFirst(elemValue)
		c.caches[k] = elemValue
	}
	if !ok {
		if len(c.caches) >= c.CacheCapacity {
			delete(c.caches, c.last.key)
			// 更新last节点
			c.removeLast()
		}
		newValue := &CacheNode{
			key:   k,
			value: v,
			pre:   nil,
			next:  nil,
		}
		c.moveToFirst(newValue)
		c.caches[k] = newValue
	}
}

func (c *LRUCache) moveToFirst(node *CacheNode) {
	if c.first == node {
		return
	}
	if node.next != nil {
		node.next.pre = node.pre
	}
	if node.pre != nil {
		node.pre.next = node.next
	}
	if node == c.last {
		c.last = node.pre
	}
	if c.first == nil || c.last == nil {
		c.first = node
		c.last = node
		return
	}
	node.next = c.first
	c.first.pre = node
	c.first = node
	c.first.pre = nil
}

func (c *LRUCache) removeLast() {
	if c.last != nil {
		c.last = c.last.pre
		if c.last == nil {
			c.first = nil
		} else {
			c.last.next = nil
		}
	}
}

func (c *LRUCache) get(k int) string {
	elem, ok := c.caches[k]
	if !ok {
		return "nil"
	}
	c.moveToFirst(elem)
	return elem.value
}

func (c *LRUCache) Print() {
	node := c.first
	for node != nil {
		fmt.Printf("%v:%v ", node.key, node.value)
		node = node.next
	}
	fmt.Println()
}

func main() {
	cache := Constructor(3)
	cache.put(1, "a")
	cache.Print()
	cache.put(2, "b")
	cache.Print()
	cache.put(3, "c")
	cache.Print()
	cache.put(4, "d")
	cache.Print()
	cache.put(1, "aa")
	cache.Print()
	cache.put(2, "bb")
	cache.Print()
	cache.put(5, "e")
	cache.Print()
	cache.get(1)
	cache.Print()
}
