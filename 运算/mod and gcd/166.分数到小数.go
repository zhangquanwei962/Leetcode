/*
 * @lc app=leetcode.cn id=166 lang=golang
 * @lcpr version=21908
 *
 * [166] 分数到小数
 */

// @lc code=start
package main

import (
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	var res strings.Builder
	if numerator < 0 != (denominator < 0) {
		res.WriteByte('-')
	}
	a, b := abs(numerator), abs(denominator)
	res.WriteString(strconv.Itoa(a / b)) // 计算整数部分
	a %= b
	if a > 0 {
		res.WriteByte('.')
	}
	m := make(map[int]int) // 记录余数出现的位置
	for a > 0 {
		if pos, ok := m[a]; ok { // 出现循环节
			res.WriteString(")")
			s := res.String()
			return s[:pos] + "(" + s[pos:]
		}
		m[a] = res.Len()
		a *= 10
		res.WriteString(strconv.Itoa(a / b))
		a %= b
	}
	return res.String()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

/*
// @lcpr case=start
// 1\n2\n
// @lcpr case=end

// @lcpr case=start
// 2\n1\n
// @lcpr case=end

// @lcpr case=start
// 4\n333\n
// @lcpr case=end

*/
