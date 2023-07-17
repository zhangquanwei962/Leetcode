/*
 * @lc app=leetcode.cn id=879 lang=golang
 * @lcpr version=21907
 *
 * [879] 盈利计划
 */

// @lc code=start
package main

func profitableSchemes(G int, P int, group []int, profit []int) int {
	sz := len(group)
	const M = 1e9 + 7
	dp := make([][]int, G+1)
	for i := range dp {
		dp[i] = make([]int, P+1)
		dp[i][0] = 1
	}
	for i := 0; i < sz; i++ {
		g, p := group[i], profit[i]
		for j := G; j >= g; j-- {
			for k := P; k >= 0; k-- {
				dp[j][k] += dp[j-g][max(k-p, 0)]
				dp[j][k] %= M
			}
		}
	}
	return dp[G][P]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// 5\n3\n[2,2]\n[2,3]\n
// @lcpr case=end

// @lcpr case=start
// 10\n5\n[2,3,5]\n[6,7,8]\n
// @lcpr case=end

*/
