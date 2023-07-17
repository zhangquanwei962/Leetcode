/*
 * @lc app=leetcode.cn id=2192 lang=golang
 * @lcpr version=21909
 *
 * [2192] 有向无环图中一个节点的所有祖先
 */

// @lc code=start

// O(VE+V^2logV)  O(V^2)
package main

import "sort"

func getAncestors(n int, edges [][]int) [][]int {
	g := make([][]int, n)
	anc := make([]map[int]bool, n) // 存储每个节点的祖先的哈希表
	indeg := make([]int, n)

	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		indeg[y]++
	}

	q := make([]int, 0)
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range g[u] {
			// 更新子节点的祖先哈希表
			if anc[v] == nil {
				anc[v] = make(map[int]bool)
			}
			anc[v][u] = true
			for i := range anc[u] {
				anc[v][i] = true
			}
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}

	// 转化为答案数组
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := range anc[i] {
			res[i] = append(res[i], j)
		}
		sort.Ints(res[i])
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 8\n[[0,3],[0,4],[1,3],[2,4],[2,7],[3,5],[3,6],[3,7],[4,6]]\n
// @lcpr case=end

// @lcpr case=start
// 5\n[[0,1],[0,2],[0,3],[0,4],[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]\n
// @lcpr case=end

*/
