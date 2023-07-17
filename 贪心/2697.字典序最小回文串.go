/*
 * @lc app=leetcode.cn id=2697 lang=golang
 * @lcpr version=21908
 *
 * [2697] 字典序最小回文串
 */

// @lc code=start
package main

func makeSmallestPalindrome(S string) string {
	s := []byte(S)
	for i, n := 0, len(s); i < n>>1; i++ {
		x, y := s[i], s[n-i-1]
		if x > y {
			s[i] = y
		} else {
			s[n-i-1] = x
		}
	}
	return string(s)
}

// @lc code=end

/*
// @lcpr case=start
// "egcfe"\n
// @lcpr case=end

// @lcpr case=start
// "abcd"\n
// @lcpr case=end

// @lcpr case=start
// "seven"\n
// @lcpr case=end

*/
