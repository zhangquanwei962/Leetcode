/*
 * @lc app=leetcode.cn id=1416 lang=golang
 * @lcpr version=21907
 *
 * [1416] 恢复数组
 */

// @lc code=start
// dp[i+1]表示以s[0...i]结尾的合法数字之和(i+1)
// 开始时可以给索引−1即0前面的初始值设为1，表示
// s[0:i+1]的合法情况
package main

import "strconv"

const MOD = 1e9 + 7

func numberOfArrays(s string, k int) int {
	n := len(s)
	m := len(strconv.Itoa(k))
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := max(i-m+1, 0); j <= i; j++ {
			if s[j] != '0' {
				num, _ := strconv.Atoi(s[j : i+1])
				if num <= k {
					dp[i+1] += dp[j]
				}
			}
		}
		dp[i+1] %= MOD
	}
	return dp[n]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

/*
// @lcpr case=start
// "1000"\n10000\n
// @lcpr case=end

// @lcpr case=start
// "1000"\n10\n
// @lcpr case=end

// @lcpr case=start
// "1317"\n2000\n
// @lcpr case=end

// @lcpr case=start
// "2020"\n30\n
// @lcpr case=end

// @lcpr case=start
// "1234567890"\n90\n
// @lcpr case=end

*/
