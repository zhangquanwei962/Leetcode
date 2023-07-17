/*
 * @lc app=leetcode.cn id=788 lang=golang
 * @lcpr version=21909
 *
 * [788] 旋转数字
 */

// @lc code=start
package main

import (
	"fmt"
	"strconv"
)

var diff = [10]int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1}

func rotatedDigits(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	dp := make([][2]int, m)
	for i := range dp {
		dp[i] = [2]int{-1, -1} // 记忆化
	}

	var f func(int, int, bool) int
	f = func(i, isDiff int, isLimit bool) (res int) {
		if i == m {
			return isDiff // 只有2/5/6/9是有效的
		}

		if !isLimit {
			dv := &dp[i][isDiff]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}

		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}

		for d := 0; d <= up; d++ {
			if diff[d] != -1 { // d不是3/4/7
				fmt.Println(isDiff, diff[d], isDiff|diff[d])
				res += f(i+1, isDiff|diff[d], isLimit && up == d)
			}
		}
		return
	}
	return f(0, 0, true)
}

// @lc code=end

/*
// @lcpr case=start
// 10\n
// @lcpr case=end

*/
