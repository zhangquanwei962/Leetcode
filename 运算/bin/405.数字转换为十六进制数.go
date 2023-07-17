/*
 * @lc app=leetcode.cn id=405 lang=golang
 * @lcpr version=21907
 *
 * [405] 数字转换为十六进制数
 */

// @lc code=start
package main

import "strings"

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	var n int64 = int64(num)
	var sb strings.Builder
	if n < 0 {
		n = (1<<32 + n)
	}
	for n != 0 {
		u := n % 16
		var c byte = byte(u + '0')
		if u >= 10 {
			c = byte(u-10) + 'a'
		}
		sb.WriteByte(c)
		n /= 16
	}
	return reverseString(sb.String())
}
func reverseString(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	b := []byte(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

// @lc code=end

/*
// @lcpr case=start
// 26\n
// @lcpr case=end

// @lcpr case=start
// -1\n
// @lcpr case=end

*/
