/*
 * @lc app=leetcode.cn id=2485 lang=golang
 * @lcpr version=21909
 *
 * [2485] 找出中枢整数
 */

// @lc code=start
package main

import "math"

func pivotInteger(n int) int {
	m := n * (n + 1) / 2
	x := int(math.Sqrt(float64(m)))
	if x*x == m {
		return int(x)
	}
	return -1
}

// @lc code=end

/*
// @lcpr case=start
// 8\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

// @lcpr case=start
// 4\n
// @lcpr case=end

*/
