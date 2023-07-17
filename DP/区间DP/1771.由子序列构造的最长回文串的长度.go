/*
 * @lc app=leetcode.cn id=1771 lang=golang
 * @lcpr version=21909
 *
 * [1771] 由子序列构造的最长回文串的长度
 */

// @lc code=start
package main

func longestPalindrome(word1, word2 string) (ans int) {
	s := word1 + word2
	n := len(s)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		f[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				f[i][j] = f[i+1][j-1] + 2
				if i < len(word1) && j >= len(word1) {
					ans = max(ans, f[i][j]) // f[i][j] 一定包含 s[i] 和 s[j]
				}
			} else {
				f[i][j] = max(f[i+1][j], f[i][j-1])
			}
		}
	}
	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// "cacb"\n"cbba"\n
// @lcpr case=end

// @lcpr case=start
// "ab"\n"ab"\n
// @lcpr case=end

// @lcpr case=start
// "aa"\n"bb"\n
// @lcpr case=end

*/
