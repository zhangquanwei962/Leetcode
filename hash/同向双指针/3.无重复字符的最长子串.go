/*
 * @lc app=leetcode.cn id=3 lang=golang
 * @lcpr version=21908
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
package main

func lengthOfLongestSubstring(s string) (ans int) {
	left := 0
	cnt := ['z' + 9]int{}

	for right, ch := range s {
		cnt[ch]++
		for cnt[ch] > 1 {
			cnt[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
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
// "abcabcbb"\n
// @lcpr case=end

// @lcpr case=start
// "bbbbb"\n
// @lcpr case=end

// @lcpr case=start
// "pwwkew"\n
// @lcpr case=end

*/
