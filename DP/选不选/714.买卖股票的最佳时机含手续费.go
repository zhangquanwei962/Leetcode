/*
 * @lc app=leetcode.cn id=714 lang=golang
 * @lcpr version=21908
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start

package main

import (
	"math"
)

func maxProfit(prices []int, fee int) int {
	f0, f1 := 0, math.MinInt/2

	for _, p := range prices {
		f0, f1 = max(f0, f1+p-fee), max(f1, f0-p)
	}
	return f0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [1, 3, 2, 8, 4, 9]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,3,7,5,10,3]\n3\n
// @lcpr case=end

*/
