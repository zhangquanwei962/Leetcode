/*
 * @lc app=leetcode.cn id=5 lang=golang
 * @lcpr version=21908
 *
 * [5] 最长回文子串
 */

// @lc code=start
package main

func longestPalindrome(s string) string {
	n := len(s)
	maxi, start1 := 0, 0
	for i := 0; i < n*2-1; i++ {
		l := i / 2
		r := i/2 + i%2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > maxi {
			maxi = r - l - 1
			start1 = l + 1
		}
	}
	return s[start1 : start1+maxi]
}

// @lc code=end

/*
// @lcpr case=start
// "babad"\n
// @lcpr case=end

// @lcpr case=start
// "cbbd"\n
// @lcpr case=end

*/
