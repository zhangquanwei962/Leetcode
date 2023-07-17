/*
 * @lc app=leetcode.cn id=2556 lang=golang
 * @lcpr version=21908
 *
 * [2556] 二进制矩阵中翻转最多一次使路径不连通
 */

// @lc code=start

// 轮廓两次遍历原地修改即可
package main

func isPossibleToCutPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int) bool

	dfs = func(i1, i2 int) bool {
		if i1 == m-1 && i2 == n-1 {
			return true
		}

		grid[i1][i2] = 0

		return i1 < m-1 && grid[i1+1][i2] > 0 && dfs(i1+1, i2) ||
			i2 < n-1 && grid[i1][i2+1] > 0 && dfs(i1, i2+1)
	}
	return !dfs(0, 0) || !dfs(0, 0)
}

// @lc code=end

/*
// @lcpr case=start
// [[1,1,1],[1,0,0],[1,1,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,1,1],[1,0,1],[1,1,1]]\n
// @lcpr case=end

*/
