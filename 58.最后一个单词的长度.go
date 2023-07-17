/*
 * @lc app=leetcode.cn id=58 lang=golang
 * @lcpr version=21908
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
package main

import "strings"

func lengthOfLastWord(s string) int {
	// 删除字符串末尾的空格
	s = strings.TrimRight(s, " ")

	// 查找最后一个单词的末尾索引
	lastSpace := strings.LastIndex(s, " ")
	if lastSpace == -1 {
		// 没有空格，整个字符串都是单词
		return len(s)
	}
	return len(s) - lastSpace - 1
}

// @lc code=end

/*
// @lcpr case=start
// "Hello World"\n
// @lcpr case=end

// @lcpr case=start
// "   fly me   to   the moon  "\n
// @lcpr case=end

// @lcpr case=start
// "luffy is still joyboy"\n
// @lcpr case=end

*/
