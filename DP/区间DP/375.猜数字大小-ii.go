/*
 * @lc app=leetcode.cn id=375 lang=golang
 * @lcpr version=21909
 *
 * [375] 猜数字大小 II
 */

// @lc code=start
package main

import "math"

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
func getMoneyAmount(n int) int {
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i+1 == j {
			return i
		}
		if i >= j {
			return 0
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		res := math.MaxInt32
		for k := i + 1; k < j; k++ {
			res = min(res, max(dfs(i, k-1), dfs(k+1, j))+k)
		}
		*p = res
		return res
	}
	return dfs(0, n)
}

// @lc code=end

/*
// @lcpr case=start
// 10\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

// @lcpr case=start
// 2\n
// @lcpr case=end

*/
