/*
 * @lc app=leetcode.cn id=310 lang=golang
 * @lcpr version=21909
 *
 * [310] 最小高度树
 */

// @lc code=start
// O(n) O(n)不算建树
// 换根dp

package main

import (
	"math"
)

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	// 建图
	graph := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// 以节点0为根，dfs遍历树，求出每个子树节点高度
	height := make([]int, n)
	var dfs func(u, fa int)
	dfs = func(u, fa int) {
		height[u] = 1
		for _, v := range graph[u] {
			if v != fa {
				dfs(v, u)
				height[u] = max(height[u], height[v]+1)
			}
		}
	}
	dfs(0, -1)
	// 换根DP，从节点0开始向其它节点转移根，过程维护树及子树高度，并记录最小高度树对应的根
	minHeight := math.MaxInt >> 1
	ans := []int{}
	var reroot func(u, fa int)
	reroot = func(u, fa int) {
		first, second := 0, 0 // 第一大，第二大
		for _, v := range graph[u] {
			if height[v] > first {
				second, first = first, height[v]
			} else if height[v] > second {
				second = height[v]
			}
		}
		// height[u] = first + 1
		if height[u] < minHeight {
			ans = []int{u}
			minHeight = height[u]
		} else if height[u] == minHeight {
			ans = append(ans, u)
		}
		for _, v := range graph[u] {
			if v != fa { // 判断u的最大子树高路径是否经过v
				if height[v] == first {
					height[u] = second + 1
				} else {
					height[u] = first + 1
				}
				height[v] = max(height[v], height[u]+1)
				reroot(v, u)
			}
		}
	}
	reroot(0, -1)

	return ans

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// func findMinHeightTrees(n int, edges [][]int) (res []int) {
// 	if n == 1 {
// 		return []int{0}
// 	}

// 	g := make([][]int, n)
// 	for _, e := range edges {
// 		x, y := e[0], e[1]
// 		g[x] = append(g[x], y)
// 		g[y] = append(g[y], x)
// 	}

// 	d1, d2, p1, up := make([]int, n), make([]int, n),
// 		make([]int, n), make([]int, n)

// 	var dfs func(int, int)
// 	dfs = func(x, fa int) {
// 		for _, y := range g[x] {
// 			if y != fa {
// 				dfs(y, x)
// 				d := d1[y] + 1
// 				if d >= d1[x] {
// 					// 先将当前最长路径给次长路径
// 					d2[x] = d1[x]
// 					// 更新最长路径
// 					d1[x] = d
// 					p1[x] = y
// 				} else if d >= d2[x] {
// 					// 只比次大值大，更新次大值
// 					d2[x] = d
// 				}
// 			}
// 		}
// 	}

// 	var reRoot func(int, int)
// 	reRoot = func(x, fa int) {
// 		for _, y := range g[x] {
// 			if y != fa {
// 				if y == p1[x] {
// 					up[y] = max(up[x], d2[x]) + 1
// 				} else {
// 					up[y] = max(up[x], d1[x]) + 1
// 				}
// 				reRoot(y, x)
// 			}
// 		}
// 	}

// 	dfs(0, -1)
// 	reRoot(0, -1)

// 	min_heigh := n + 1
// 	for i := 0; i < n; i++ {
// 		// 遍历所有节点，找到最小的高度
// 		min_heigh = min(min_heigh, max(d1[i], up[i]))
// 	}
// 	for i := 0; i < n; i++ {
// 		// 再次遍历所有节点，如果该点高度等于最小高度则加入答案， 这一步也可以合并到上一步中，为了清晰则不单独写。
// 		if max(d1[i], up[i]) == min_heigh {
// 			res = append(res, i)
// 		}
// 	}
// 	return
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }

// func findMinHeightTrees(n int, edges [][]int) []int {
// 	if n == 1 {
// 		return []int{0}
// 	}

// 	g := make([][]int, n)
// 	for _, e := range edges {
// 		x, y := e[0], e[1]
// 		g[x] = append(g[x], y)
// 		g[y] = append(g[y], x)
// 	}

// 	parents := make([]int, n)
// 	maxDepth, node := 0, -1
// 	var dfs func(int, int, int)
// 	dfs = func(x, pa, depth int) {
// 		if depth > maxDepth {
// 			maxDepth, node = depth, x
// 		}
// 		parents[x] = pa
// 		for _, y := range g[x] {
// 			if y != pa {
// 				dfs(y, x, depth+1)
// 			}
// 		}
// 	}
// 	dfs(0, -1, 1)
// 	maxDepth = 0
// 	dfs(node, -1, 1)

// 	path := []int{}
// 	for node != -1 {
// 		path = append(path, node)
// 		node = parents[node]
// 	}
// 	m := len(path)
// 	if m%2 == 0 {
// 		return []int{path[m/2-1], path[m/2]}
// 	}
// 	return []int{path[m/2]}
// }

// @lc code=end

/*
// @lcpr case=start
// 4\n[[1,0],[1,2],[1,3]]\n
// @lcpr case=end

// @lcpr case=start
// 6\n[[3,0],[3,1],[3,2],[3,4],[5,4]]\n
// @lcpr case=end

*/
