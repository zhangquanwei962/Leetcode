/*
 * @lc app=leetcode.cn id=1642 lang=golang
 * @lcpr version=21909
 *
 * [1642] 可以到达的最远建筑
 */

// @lc code=start
// 可以预见，贪心策略，
// 对于梯子，肯定钢刀用到刀刃上，维护k个最大delta h

package main

import (
	"container/heap"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	n := len(heights)
	q := &MinHeap{}
	// heap.Init(q)
	sumH := 0
	for i := 1; i < n; i++ {
		deltaH := heights[i] - heights[i-1]
		if deltaH > 0 {
			heap.Push(q, deltaH)
			if q.Len() > ladders {
				sumH += heap.Pop(q).(int)
			}
			if sumH > bricks {
				return i - 1
			}
		}
	}
	return n - 1
}

// @lc code=end

/*
// @lcpr case=start
// [4,2,7,6,9,14,12]\n5\n1\n
// @lcpr case=end

// @lcpr case=start
// [4,12,2,7,3,18,20,3,19]\n10\n2\n
// @lcpr case=end

// @lcpr case=start
// [14,3,19,3]\n17\n0\n
// @lcpr case=end

*/
