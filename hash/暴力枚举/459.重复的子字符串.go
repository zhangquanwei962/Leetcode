/*
 * @lc app=leetcode.cn id=459 lang=golang
 * @lcpr version=21909
 *
 * [459] 重复的子字符串
 */

// @lc code=start
package main

import (
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	for i := n / 2; i >= 1; i-- {
		if n%i == 0 && strings.Repeat(s[:i], n/i) == s {
			return true
		}
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// "abab"\n
// @lcpr case=end

// @lcpr case=start
// "aba"\n
// @lcpr case=end

// @lcpr case=start
// "abcabcabcabc"\n
// @lcpr case=end

*/
