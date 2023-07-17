/*
 * @lc app=leetcode.cn id=2384 lang=golang
 * @lcpr version=21908
 *
 * [2384] 最大回文数字
 */

// @lc code=start
package main

import (
	"fmt"
	"strings"
)

func largestPalindromic(num string) string {
	cnt := [byte('9') + 1]int{}
	for _, d := range num {
		cnt[d]++
	}

	if cnt['0'] == len(num) { // 全0
		return "0"
	}

	s := []byte{}
	for i := byte('9'); i > '0' || i == '0' && len(s) > 0; i-- {
		s = append(s, strings.Repeat(string(i), cnt[i]/2)...)
	}

	j := len(s) - 1
	for i := '9'; i >= '0'; i-- {
		if cnt[i]&1 > 0 { // 还可以填一个变成奇回文串
			s = append(s, byte(i))
			break
		}
	}
	fmt.Println('9' == rune('9'))
	for ; j >= 0; j-- { // 添加镜像部分
		s = append(s, s[j])
	}
	return string(s)
}

// @lc code=end

/*
// @lcpr case=start
// "444947137"\n
// @lcpr case=end

// @lcpr case=start
// "00009"\n
// @lcpr case=end

*/
