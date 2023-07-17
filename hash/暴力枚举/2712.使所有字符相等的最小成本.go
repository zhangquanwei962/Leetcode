/*
 * @lc app=leetcode.cn id=2712 lang=golang
 * @lcpr version=21909
 *
 * [2712] 使所有字符相等的最小成本
 */

// @lc code=start
package main

func minimumCost(s string) (ans int64) {
	n := len(s)
	for i := 1; i < n; i++ {
		if s[i-1] != s[i] {
			ans += int64(min(i, n-i))
		}
	}
	return
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// "0011"\n
// @lcpr case=end

// @lcpr case=start
// "010101"\n
// @lcpr case=end

*/
