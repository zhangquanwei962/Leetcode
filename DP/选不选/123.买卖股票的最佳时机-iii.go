/*
 * @lc app=leetcode.cn id=123 lang=golang
 * @lcpr version=21908
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
package main

import "math"

func maxProfit(prices []int) int {
	k := 2
	f := make([][2]int, k+2)

	for j := 1; j <= k+1; j++ {
		f[j][1] = math.MinInt / 2
	}

	f[0][0] = math.MinInt / 2

	for _, p := range prices {
		for j := k + 1; j > 0; j-- {
			f[j][0] = max(f[j][0], f[j][1]+p)
			f[j][1] = max(f[j][1], f[j-1][0]-p)
		}
	}
	return f[k+1][0]
}

// func maxProfit(k int, prices []int) int {
// 	n := len(prices)

// 	f := make([][][2]int, n+1)
// 	for i := range f {
// 		f[i] = make([][2]int, k+2)
// 		for j := range f[i] {
// 			f[i][j] = [2]int{math.MinInt / 2, math.MinInt / 2}
// 		}
// 	}

// 	for j := 1; j <= k+1; j++ {
// 		f[0][j][0] = 0
// 	}

// 	for i, p := range prices {
// 		for j := 1; j <= k+1; j++ {
// 			f[i+1][j][0] = max(f[i][j][0], f[i][j][1]+p)
// 			f[i+1][j][1] = max(f[i][j][1], f[i][j-1][0]-p)
// 		}
// 	}
// 	return f[n][k+1][0]
// }

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [3,3,5,0,0,3,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,4,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/
