/*
 * @lc app=leetcode.cn id=1631 lang=golang
 * @lcpr version=21909
 *
 * [1631] 最小体力消耗路径
 */

// @lc code=start
package main

import (
	"container/heap"
	"math"
)

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt
		}
	}
	var dirs4 = [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	dis[0][0] = 0
	h := &hp{{}}

	for {
		p := heap.Pop(h).(tuple)
		d, i, j := p.d, p.i, p.j
		if d > dis[i][j] {
			continue
		}

		if i == m-1 && j == n-1 { // 找到终点，此时 d 一定是最短路
			return d
		}

		for _, v := range dirs4 {
			x, y := i+v[0], j+v[1]
			if 0 <= x && x < m && 0 <= y && y < n {
				nd := max(d, abs(heights[i][j]-heights[x][y]))
				if nd < dis[x][y] {
					dis[x][y] = nd
					heap.Push(h, tuple{nd, x, y})
				}
			}
		}
	}
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

type tuple struct{ d, i, j int }
type hp []tuple

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,2],[3,8,2],[5,3,5]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[3,8,4],[5,3,5]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]\n
// @lcpr case=end

*/
