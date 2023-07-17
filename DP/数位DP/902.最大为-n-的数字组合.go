/*
 * @lc app=leetcode.cn id=902 lang=golang
 * @lcpr version=21909
 *
 * [902] 最大为 N 的数字组合
 */

// @lc code=start
// O(d logn) O(logn)
package main

import (
	"strconv"
)

func atMostNGivenDigitSet(digits []string, n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	dp := make([]int, m)
	for i := range dp {
		dp[i] = -1 // 记忆化
	}

	var f func(int, bool, bool) int
	f = func(i int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum { // 填入了数字
				return 1
			}
			return
		}

		if !isLimit && isNum {
			dv := &dp[i]
			if *dv >= 0 {
				return *dv
			}
			defer func() {
				*dv = res
			}()
		}

		if !isNum { // 可以跳过当前数位
			res += f(i+1, false, false)
		}

		up := byte('9')
		if isLimit {
			up = s[i]
		}

		for _, d := range digits {
			if d[0] > up { // 如果d[0]>up 超了
				break
			}
			res += f(i+1, isLimit && up == d[0], true)
		}
		return
	}
	return f(0, true, false)
}

// @lc code=end

/*
// @lcpr case=start
// ["1","3","5","7"]\n100\n
// @lcpr case=end

// @lcpr case=start
// ["1","4","9"]\n1000000000\n
// @lcpr case=end

// @lcpr case=start
// ["7"]\n8\n
// @lcpr case=end

*/
