/*
 * @lc app=leetcode.cn id=120 lang=golang
 * @lcpr version=21907
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
package main

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func minimumTotal(triangle [][]int) int {
	m, n := len(triangle), len(triangle)

	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没被计算过
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == m {
			return 0
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		*p = min(dfs(i+1, j), dfs(i+1, j+1)) + triangle[i][j]
		return *p
	}
	return dfs(0, 0)
}

// @lc code=end

/*
// @lcpr case=start
// [[2],[3,4],[6,5,7],[4,1,8,3]]\n
// @lcpr case=end

// @lcpr case=start
// [[-10]]\n
// @lcpr case=end

*/
