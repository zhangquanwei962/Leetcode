/*
 * @lc app=leetcode.cn id=600 lang=golang
 * @lcpr version=21909
 *
 * [600] 不含连续1的非负整数
 */

// @lc code=start
package main

import "strconv"

func findIntegers(n int) int {
	s := strconv.FormatInt(int64(n), 2)
	m := len(s)
	dp := make([][2]int, m)
	for i := range dp {
		dp[i] = [2]int{-1, -1}
	}
	var f func(int, int8, bool) int
	f = func(i int, pre int8, isLimit bool) (res int) {
		if i == m {
			return 1
		}
		if !isLimit {
			dv := &dp[i][pre]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}

		up := 1
		if isLimit {
			up = int(s[i] & 1) // 有限制
		}
		// 如果在受到约束的情况下填了 s[i]，那么后续填入的数字仍会受到 n 的约束
		res = f(i+1, 0, isLimit && up == 0) // 选0
		if pre == 0 && up == 1 {
			res += f(i+1, 1, isLimit) // 选1
		}
		return
	}
	return f(0, 0, true)
}

// @lc code=end

/*
// @lcpr case=start
// 5\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

// @lcpr case=start
// 2\n
// @lcpr case=end

*/
