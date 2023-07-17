/*
 * @lc app=leetcode.cn id=121 lang=golang
 * @lcpr version=21908
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
package main

func maxProfit(prices []int) (ans int) {
	minprice, n := prices[0], len(prices)
	for i := 1; i < n; i++ {
		ans = max(ans, prices[i]-minprice)
		minprice = min(minprice, prices[i])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [7,1,5,3,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,4,3,1]\n
// @lcpr case=end

*/
