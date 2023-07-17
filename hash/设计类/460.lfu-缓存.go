/*
 * @lc app=leetcode.cn id=460 lang=golang
 * @lcpr version=21909
 *
 * [460] LFU 缓存
 */

// @lc code=start
package main

import "container/heap"

type LFUCache struct {
	h         *NodeHeap
	nodemap   map[int]*Node
	size, cap int
	opcnt     int
}

type Node struct {
	freq, lastop int
	key, value   int
	index        int
}
type NodeHeap []*Node

func (h NodeHeap) Len() int { return len(h) }
func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index, h[j].index = h[j].index, h[i].index
}
func (h NodeHeap) Less(i, j int) bool {
	if h[i].freq < h[j].freq {
		return true
	} else if h[i].freq > h[j].freq {
		return false
	} else {
		return h[i].lastop < h[j].lastop
	}
}
func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}
func (h *NodeHeap) Pop() interface{} {
	ret := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return ret
}

func Constructor(capacity int) LFUCache {
	cache := LFUCache{&NodeHeap{},
		map[int]*Node{},
		0,
		capacity,
		0}
	heap.Init(cache.h)
	return cache
}

func (this *LFUCache) Get(key int) int {
	if this.cap == 0 {
		return -1
	}
	node, ok := this.nodemap[key]
	if !ok {
		return -1
	} else {
		ret := node.value
		node.freq++
		node.lastop = this.opcnt
		this.opcnt++
		heap.Fix(this.h, node.index)
		return ret
	}
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}
	node, ok := this.nodemap[key]
	if ok {
		node.freq++
		node.lastop = this.opcnt
		this.opcnt++
		node.value = value
		heap.Fix(this.h, node.index)
	} else {
		if this.size == this.cap {
			kicked := heap.Pop(this.h)
			delete(this.nodemap, kicked.(*Node).key)
		} else {
			this.size++
		}
		newnode := &Node{1, this.opcnt, key, value, this.h.Len()}
		this.opcnt++
		heap.Push(this.h, newnode)
		this.nodemap[key] = newnode
	}
}

// type LFUCache struct {
// }

// func Constructor(capacity int) LFUCache {

// }

// func (this *LFUCache) Get(key int) int {

// }

// func (this *LFUCache) Put(key int, value int) {

// }

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

/*
// @lcpr case=start
// ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"][[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]\n
// @lcpr case=end

*/
