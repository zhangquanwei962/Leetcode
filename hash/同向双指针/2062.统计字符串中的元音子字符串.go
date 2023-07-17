/*
 * @lc app=leetcode.cn id=2062 lang=golang
 * @lcpr version=21908
 *
 * [2062] 统计字符串中的元音子字符串
 */

// @lc code=start
// 双指针
package main

import (
	"strings"
)

func countVowelSubstrings(word string) (ans int) {
	for _, s := range strings.FieldsFunc(word, func(r rune) bool { return !strings.ContainsRune("aeiou", r) }) { // 分割出仅包含元音的字符串
		cnt := ['v']int{}
		left := 0
		for _, ch := range s {
			cnt[ch]++
			for cnt[s[left]] > 1 { // 双指针，仅当该元音个数不止一个时才移动左指针
				cnt[s[left]]--
				left++
			}
			if cnt['a'] > 0 && cnt['e'] > 0 && cnt['i'] > 0 && cnt['o'] > 0 && cnt['u'] > 0 { // 必须包含全部五种元音
				ans += left + 1
			}
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// "aeiouu"\n
// @lcpr case=end

// @lcpr case=start
// "unicornarihan"\n
// @lcpr case=end

// @lcpr case=start
// "cuaieuouac"\n
// @lcpr case=end

// @lcpr case=start
// "bbaeixoubb"\n
// @lcpr case=end

*/
