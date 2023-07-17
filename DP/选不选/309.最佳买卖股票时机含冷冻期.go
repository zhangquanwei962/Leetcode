/*
 * @lc app=leetcode.cn id=309 lang=golang
 * @lcpr version=21908
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
package main

import "math"

func maxProfit(prices []int) int {
	pre0, f0, f1 := 0, 0, math.MinInt
	for _, p := range prices {
		pre0, f0, f1 = f0, max(f0, f1+p), max(f1, pre0-p)
	}
	return f0
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,0,2]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/
