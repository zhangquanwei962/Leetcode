/*
 * @lc app=leetcode.cn id=2063 lang=golang
 * @lcpr version=21908
 *
 * [2063] 所有子字符串中的元音
 */

// @lc code=start
// 贡献法
package main

import "strings"

func countVowels(word string) (ans int64) {
	for i, ch := range word {
		if strings.ContainsRune("aeiou", ch) {
			ans += int64(i+1) * int64(len(word)-i)
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// "aba"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n
// @lcpr case=end

// @lcpr case=start
// "ltcd"\n
// @lcpr case=end

// @lcpr case=start
// "noosabasboosa"\n
// @lcpr case=end

*/
