/*
 * @lc app=leetcode.cn id=743 lang=golang
 * @lcpr version=21909
 *
 * [743] 网络延迟时间
 */

// @lc code=start
package main

import (
	"container/heap"
	"math"
)

func networkDelayTime(times [][]int, n int, k int) (ans int) {
	g := make([][]neighbor, n)
	for _, e := range times {
		x, y, w := e[0]-1, e[1]-1, e[2]
		g[x] = append(g[x], neighbor{y, w})
	}
	dist := dijkstra(g, k-1)
	for _, d := range dist {
		if d == math.MaxInt32 {
			return -1
		}
		ans = max(ans, d)
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func networkDelayTime(times [][]int, n, k int) (ans int) {
//     const inf = math.MaxInt64 / 2
//     g := make([][]int, n)
//     for i := range g {
//         g[i] = make([]int, n)
//         for j := range g[i] {
//             g[i][j] = inf
//         }
//     }
//     for _, t := range times {
//         x, y := t[0]-1, t[1]-1
//         g[x][y] = t[2]
//     }

//     dist := make([]int, n)
//     for i := range dist {
//         dist[i] = inf
//     }
//     dist[k-1] = 0
//     used := make([]bool, n)
//     for i := 0; i < n; i++ {
//         x := -1
//         for y, u := range used {
//             if !u && (x == -1 || dist[y] < dist[x]) {
//                 x = y
//             }
//         }
//         used[x] = true
//         for y, time := range g[x] {
//             dist[y] = min(dist[y], dist[x]+time)
//         }
//     }

//     for _, d := range dist {
//         if d == inf {
//             return -1
//         }
//         ans = max(ans, d)
//     }
//     return
// }

// func min(a, b int) int {
//     if a < b {
//         return a
//     }
//     return b
// }

// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }

// @lc code=end

/*
// @lcpr case=start
// [[2,1,1],[2,3,1],[3,4,1]]\n4\n2\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,1]]\n2\n1\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,1]]\n2\n2\n
// @lcpr case=end

*/
