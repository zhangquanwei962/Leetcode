/*
 * @lc app=leetcode.cn id=882 lang=golang
 * @lcpr version=21909
 *
 * [882] 细分图中的可到达节点
 */

// @lc code=start
// O(n+mlogm) O(n+m)
package main

import (
	"container/heap"
	"math"
)

func reachableNodes(edges [][]int, maxMoves, n int) (ans int) {
	g := make([][]neighbor, n)
	for _, e := range edges {
		u, v, cnt := e[0], e[1], e[2]
		g[u] = append(g[u], neighbor{v, cnt + 1})
		g[v] = append(g[v], neighbor{u, cnt + 1}) // 建图
	}

	dist := dijkstra(g, 0) // 从 0 出发的最短路

	for _, d := range dist {
		if d <= maxMoves { // 这个点可以在 maxMoves 步内到达
			ans++
		}
	}
	for _, e := range edges {
		u, v, cnt := e[0], e[1], e[2]
		a := max(maxMoves-dist[u], 0)
		b := max(maxMoves-dist[v], 0)
		ans += min(a+b, cnt) // 这条边上可以到达的节点数
	}
	return
}

// 以下为 Dijkstra 算法模板
type neighbor struct{ to, weight int }

func dijkstra(g [][]neighbor, start int) []int {
	dist := make([]int, len(g))
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0
	h := hp{{start, 0}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		x := p.x
		if dist[x] < p.dist {
			continue
		}
		for _, e := range g[x] {
			y := e.to
			if d := dist[x] + e.weight; d < dist[y] {
				dist[y] = d
				heap.Push(&h, pair{y, d})
			}
		}
	}
	return dist
}

type pair struct{ x, dist int }
type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].dist < h[j].dist }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [[0,1,10],[0,2,1],[1,2,2]]\n6\n3\n
// @lcpr case=end

// @lcpr case=start
// [[0,1,4],[1,2,6],[0,2,8],[1,3,1]]\n10\n4\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,4],[1,4,5],[1,3,1],[2,3,4],[3,4,5]]\n17\n5\n
// @lcpr case=end

*/
