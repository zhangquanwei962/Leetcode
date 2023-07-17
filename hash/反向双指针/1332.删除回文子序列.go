/*
 * @lc app=leetcode.cn id=1332 lang=golang
 * @lcpr version=21909
 *
 * [1332] 删除回文子序列
 */

// @lc code=start
package main

func removePalindromeSub(s string) int {
	for i, n := 0, len(s); i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return 2
		}
	}
	return 1
}

// @lc code=end

/*
// @lcpr case=start
// "ababa"\n
// @lcpr case=end

// @lcpr case=start
// "abb"\n
// @lcpr case=end

// @lcpr case=start
// "baabb"\n
// @lcpr case=end

*/
