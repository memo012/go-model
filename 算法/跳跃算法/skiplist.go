package main

import (
	"math/rand"
)

type Node struct {
	value int
	right *Node
	down  *Node
}

type Skiplist struct {
	head *Node
}

func Constructor() Skiplist {
	return Skiplist{head: nil}
}

var level = 0

func (this *Skiplist) Search(target int) bool {
	node := this.head
	for node != nil {
		for node.right != nil && node.right.value < target {
			node = node.right
		}
		if node.right != nil && node.right.value == target {
			return true
		}
		node = node.down
	}
	return false
}

func (this *Skiplist) Add(num int) {
	rlevel := 1
	for rlevel <= level && throwCoin() {
		rlevel++
	}
	if rlevel > level {
		level = rlevel
		this.head = &Node{
			value: num,
			right: nil,
			down: this.head,
		}
	}
	current := this.head
	var last *Node = nil
	for i  := level; i >= 1; i-- {
		for current.right != nil && current.right.value < num {
			current = current.right
		}
		if i <= rlevel {
			current.right = &Node{
				value: num,
				right: current.right,
				down: nil,
			}
			if last != nil {
				last.down = current.right
			}
			last = current.right
		}
		current = current.down
	}
}

func (this *Skiplist) Erase(num int) bool {
	node := this.head
	var seen bool = false
	for i := level; i >= 1; i-- {
		for node.right != nil && node.right.value < num {
			node = node.right
		}
		if node.right != nil && node.right.value == num {
			seen = true
			tmp := node.right
			node.right = tmp.right
		}
		node = node.down
	}
	return seen
}

func throwCoin() bool {
	return rand.Float64() < 0.5
}

func main() {
	skiplist := Constructor()
	skiplist.Add(5)
	skiplist.Add(3)
}


