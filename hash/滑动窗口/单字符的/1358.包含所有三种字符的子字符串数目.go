/*
 * @lc app=leetcode.cn id=1358 lang=golang
 * @lcpr version=21908
 *
 * [1358] 包含所有三种字符的子字符串数目
 */

// @lc code=start
package main

import "strings"

func numberOfSubstrings(S string) (ans int) {
	for _, s := range strings.FieldsFunc(S, func(r rune) bool { return !strings.ContainsRune("abc", r) }) {
		cnt := ['c' + 1]int{}
		left := 0
		for _, ch := range s {
			cnt[ch]++
			for cnt[s[left]] > 1 {
				cnt[s[left]]--
				left++
			}

			if cnt['a'] > 0 && cnt['b'] > 0 && cnt['c'] > 0 {
				ans += left + 1
			}
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// "abcabc"\n
// @lcpr case=end

// @lcpr case=start
// "aaacb"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n
// @lcpr case=end

*/
