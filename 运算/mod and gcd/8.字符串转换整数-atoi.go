/*
 * @lc app=leetcode.cn id=8 lang=golang
 * @lcpr version=21909
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
package main

import (
	"math"
	"unicode"
)

func myAtoi(s string) (res int) {
	sign := 1
	n := len(s)
	i := 0
	for i < n && s[i] == ' ' {
		i++
	}
	start := i
	for ; i < n; i++ {
		c := s[i]
		if i == start && c == '+' {
			sign = 1
		} else if i == start && c == '-' {
			sign = -1
		} else if unicode.IsDigit(rune(c)) {
			num := int(c - '0')
			if res < math.MinInt32/10 || res == math.MinInt32/10 && -num < math.MinInt32%10 {
				return math.MinInt32
			} else if res > math.MaxInt32/10 || res == math.MaxInt32/10 && num > math.MaxInt32%10 {
				return math.MaxInt32
			}
			res = res*10 + sign*num
		} else {
			break
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// "42"\n
// @lcpr case=end

// @lcpr case=start
// "   -42"\n
// @lcpr case=end

// @lcpr case=start
// "4193 with words"\n
// @lcpr case=end

*/
