/*
 * @lc app=leetcode.cn id=214 lang=golang
 * @lcpr version=21909
 *
 * [214] 最短回文串
 */

// @lc code=start
package main

func shortestPalindrome(s string) string {
	p, n := s[:], len(s)
	if n == 0 {
		return ""
	}
	ne := make([]int, n)
	ne[0] = -1

	for i, j := 1, -1; i < n; i++ {
		for j > -1 && p[i] != p[j+1] {
			j = ne[j]
		}
		if p[i] == p[j+1] {
			j++
		}
		ne[i] = j
	}

	j := -1

	for i := n - 1; i >= 0; i-- {
		for j > -1 && s[i] != p[j+1] {
			j = ne[j]
		}
		if s[i] == p[j+1] {
			j++

		}
		// if j == n-1 {
		// 	j = ne[n-1]
		// }
	}
	add := ""

	if j != n-1 {
		add = s[j+1:]
	}
	b := []byte(add)
	for i := 0; i < len(b)>>1; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return string(b) + s
}

// @lc code=end

/*
// @lcpr case=start
// "aacecaaa"\n
// @lcpr case=end

// @lcpr case=start
// "abcd"\n
// @lcpr case=end

*/
