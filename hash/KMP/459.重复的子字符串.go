/*
 * @lc app=leetcode.cn id=459 lang=golang
 * @lcpr version=21909
 *
 * [459] 重复的子字符串
 */

// @lc code=start
// O(n) O(n)
package main

func repeatedSubstringPattern(s string) bool {
	p := s
	s = s + s
	n, m := len(s), len(p)
	ne := make([]int, m)
	ne[0] = -1

	for i, j := 1, -1; i < m; i++ {
		for j > -1 && p[i] != p[j+1] {
			j = ne[j]
		}
		if p[i] == p[j+1] {
			j++
		}
		ne[i] = j
	}

	for i, j := 1, -1; i < n-1; i++ {
		for j > -1 && s[i] != p[j+1] {
			j = ne[j]
		}
		if s[i] == p[j+1] {
			j++

		}
		if j == m-1 {
			return true
		}
	}
	return false
}

// package main

// import (
// 	"strings"
// )

// func repeatedSubstringPattern(s string) bool {
// 	n := len(s)
// 	for i := n / 2; i >= 1; i-- {
// 		if n%i == 0 && strings.Repeat(s[:i], n/i) == s {
// 			return true
// 		}
// 	}
// 	return false
// }

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
