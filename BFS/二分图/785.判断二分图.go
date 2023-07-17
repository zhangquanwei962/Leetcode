/*
 * @lc app=leetcode.cn id=785 lang=golang
 * @lcpr version=21909
 *
 * [785] 判断二分图
 */

// @lc code=start
package main

func isBipartite(graph [][]int) bool {
	n := len(graph)
	colors := make([]int8, n)
	var nodes []int
	var dfs func(int, int8) bool
	dfs = func(x int, c int8) bool {
		nodes = append(nodes, x)
		colors[x] = c
		for _, y := range graph[x] {
			if colors[y] == c || colors[y] == 0 && !dfs(y, -c) {
				return false
			}
		}
		return true
	}

	for i, c := range colors {
		if c == 0 {
			nodes = nil
			if !dfs(i, 1) { // 如果不是二分图（有奇环），则无法分组
				return false
			}
		}
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[0,2],[0,1,3],[0,2]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,3],[0,2],[1,3],[0,2]]\n
// @lcpr case=end

*/
