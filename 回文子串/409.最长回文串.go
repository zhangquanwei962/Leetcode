/*
 * @lc app=leetcode.cn id=409 lang=golang
 * @lcpr version=21909
 *
 * [409] 最长回文串
 */

// @lc code=start
package main

func longestPalindrome(s string) (ans int) {
	count := make(map[rune]int)
	for _, c := range s {
		count[c]++
	}

	for _, v := range count {
		ans += v / 2 * 2
		if v&1 == 1 && ans&1 == 0 {
			ans++
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// "abccccdd"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n
// @lcpr case=end

// @lcpr case=start
// "aaaaaccc"\n
// @lcpr case=end

*/
