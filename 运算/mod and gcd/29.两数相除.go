/*
 * @lc app=leetcode.cn id=29 lang=golang
 * @lcpr version=21908
 *
 * [29] 两数相除
 */

// @lc code=start

// O(logn)
package main

import "math"

func divide(dividend int, divisor int) (ans int) {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	positive := true
	if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
		positive = false
	}
	a, b := abs(dividend), abs(divisor)

	for i := 31; i >= 0; i-- {
		if (a >> i) >= b {
			ans += 1 << i
			a -= b << i
		}
	}
	if !positive {
		ans = -ans
	}
	return
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// @lc code=end

/*
// @lcpr case=start
// 10\n3\n
// @lcpr case=end

// @lcpr case=start
// 7\n-3\n
// @lcpr case=end

*/
