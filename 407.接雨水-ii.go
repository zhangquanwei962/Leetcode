/*
 * @lc app=leetcode.cn id=407 lang=golang
 * @lcpr version=21909
 *
 * [407] 接雨水 II
 */

// @lc code=start
// 本质是为了求四周最低接雨水高度
// 每个格子是一个点，相邻两格子之间有一条边
// 每条边的权重为两端点高度的最大值

//先「从(x,y) 出发到达边界的路径」也可看作「从边界到达点
// (x,y) 的路径」，经过转换操作后，直接计算边界到点
// 	(x,y) 的路径是一个多源点问题
// 我们可以考虑引入「超级源点」进行简化：超级源点与所有的边界格子连有一条权值为
// 0 的边，从而进一步将问题转化为「求从超级源点出发到达 (x,y) 的路径高度的最小值」。
package main

import "container/heap"

func trapRainWater(heightMap [][]int) (ans int) {
	m, n := len(heightMap), len(heightMap[0])
	if m <= 2 || n <= 2 {
		return 0
	}

	var dirs4 = [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	h := &hp{{}}
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	for i := 1; i < n; i++ {
		heap.Push(h, tuple{heightMap[0][i], 0, i})
		heap.Push(h, tuple{heightMap[m-1][i], m - 1, i})
	}

	for i := 1; i < m; i++ {
		heap.Push(h, tuple{heightMap[i][0], i, 0})
		heap.Push(h, tuple{heightMap[i][n-1], i, n - 1})
	}

	for len(*h) > 0 {
		p := heap.Pop(h).(tuple)
		d, i, j := p.d, p.i, p.j
		for _, v := range dirs4 {
			x, y := i+v[0], j+v[1]
			if 0 < x && x < m-1 && 0 < y && y < n-1 && !vis[x][y] {
				vis[x][y] = true
				new_d := max(d, heightMap[x][y]) // 接水后高度
				ans += max(0, d-heightMap[x][y]) // 实际接水高度
				heap.Push(h, tuple{new_d, x, y})
			}
		}
	}
	return
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
// [[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[3,3,3,3,3],[3,2,2,2,3],[3,2,1,2,3],[3,2,2,2,3],[3,3,3,3,3]]\n
// @lcpr case=end

*/
