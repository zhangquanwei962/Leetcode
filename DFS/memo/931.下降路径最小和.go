/*
 * @lc app=leetcode.cn id=931 lang=golang
 * @lcpr version=21907
 *
 * [931] 下降路径最小和
 */

// @lc code=start
package main

import "math"

func minFallingPathSum(matrix [][]int) int {

	m, n := len(matrix), len(matrix[0])

	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没被计算过
		}
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if j < 0 || j >= n {
			return math.MaxInt
		}
		if i == 0 {
			return matrix[0][j]
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		*p = min(dfs(i-1, j), min(dfs(i-1, j+1), dfs(i-1, j-1))) + matrix[i][j]
		return *p
	}
	res := math.MaxInt
	for i := 0; i < n; i++ {
		res = min(res, dfs(m-1, i))
	}
	return res

}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end

/*
// @lcpr case=start
// [[2,1,3],[6,5,4],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[-19,57],[-40,-5]]\n
// @lcpr case=end

*/
