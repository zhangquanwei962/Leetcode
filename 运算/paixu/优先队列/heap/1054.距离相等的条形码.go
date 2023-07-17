/*
 * @lc app=leetcode.cn id=1054 lang=golang
 * @lcpr version=21907
 *
 * [1054] 距离相等的条形码
 */

// @lc code=start

// O(logn)
package main

import "container/heap"

type PriorityQueue [][]int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i][0] > pq[j][0]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.([]int)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func rearrangeBarcodes(barcodes []int) []int {
	cout := make(map[int]int)
	for _, b := range barcodes {
		cout[b]++
	}

	q := &PriorityQueue{}
	heap.Init(q)
	for k, v := range cout {
		heap.Push(q, []int{v, k})
	}
	n := len(barcodes)
	res := make([]int, n)

	for i := 0; i < n; i++ {
		p := heap.Pop(q).([]int)
		cx, x := p[0], p[1]
		if i == 0 || res[i-1] != x {
			res[i] = x
			if cx > 1 {
				heap.Push(q, []int{cx - 1, x})
			}
		} else {
			p2 := heap.Pop(q).([]int)
			cy, y := p2[0], p2[1]
			res[i] = y
			if cy > 1 {
				heap.Push(q, []int{cy - 1, y})
			}
			heap.Push(q, []int{cx, x})
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,1,2,2,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,1,2,2,3,3]\n
// @lcpr case=end

*/
