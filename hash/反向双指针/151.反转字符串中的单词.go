/*
 * @lc app=leetcode.cn id=151 lang=golang
 * @lcpr version=21909
 *
 * [151] 反转字符串中的单词
 */

// @lc code=start
package main

import "strings"

func reverseWords(s string) string {
	words := strings.Fields(s)
	for i, n := 0, len(words); i < n>>1; i++ {
		words[i], words[n-i-1] = words[n-i-1], words[i]
	}
	return strings.Join(words, " ")
}

// @lc code=end

/*
// @lcpr case=start
// "the sky is blue"\n
// @lcpr case=end

// @lcpr case=start
// "  hello world  "\n
// @lcpr case=end

// @lcpr case=start
// "a good   example"\n
// @lcpr case=end

*/
