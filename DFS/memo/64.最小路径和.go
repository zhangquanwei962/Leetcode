/*
 * @lc app=leetcode.cn id=64 lang=golang
 * @lcpr version=21907
 *
 * [64] 最小路径和
 */

// @lc code=start
package main

import "math"

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没被计算过
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i < 0 || j < 0 {
			return math.MaxInt
		}
		if j == 0 && i == 0 {
			return grid[0][0]
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		res = min(dfs(i-1, j), dfs(i, j-1)) + grid[i][j]
		return res
	}
	return dfs(m-1, n-1)
}

// @lc code=end

/*
// @lcpr case=start
// [[1,3,1],[1,5,1],[4,2,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[4,5,6]]\n
// @lcpr case=end

*/
