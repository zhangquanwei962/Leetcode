/*
 * @lc app=leetcode.cn id=122 lang=golang
 * @lcpr version=21908
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start

// dfs(i,0)表示到第i天结束的时候，未持有股票的最大利润
// dfs(i,1)表示到第i天结束时候，持有股票的最大利润
// 由于第i-1天的结束也是第i天的开始，dfs(i-1,>)也表示第
// i天开始的最大利润
package main

import "math"

func maxProfit(prices []int) int {
	f0, f1 := 0, math.MinInt // f0未持有股票 f1持有股票
	for _, p := range prices {
		f0, f1 = max(f0, f1+p), max(f1, f0-p)
	}
	return f0
}

// func maxProfit(prices []int) int {
// 	f0, f1 := 0, math.MinInt
// 	for _, p := range prices {
// 		new_f0 := max(f0, f1+p)
// 		f1 = max(f1, f0-p)
// 		f0 = new_f0
// 	}
// 	return f0
// }

// func maxProfit(prices []int) int {
// 	n := len(prices)
// 	f := make([][2]int, n+1)
// 	f[0][1] = math.MinInt

// 	for i, p := range prices {
// 		f[i+1][0] = max(f[i][0], f[i][1]+p)
// 		f[i+1][1] = max(f[i][1], f[i][0]-p)
// 	}
// 	return f[n][0]
// }

// func maxProfit(prices []int) int {
// 	n := len(prices)
// 	memo := make([][2]int, n)
// 	for i := range memo {
// 		memo[i] = [2]int{-1, -1}
// 	}

// 	var dfs func(int, int) int
// 	dfs = func(i, hold int) (res int) {
// 		if i < 0 {
// 			if hold == 1 {
// 				return math.MinInt
// 			}
// 			return
// 		}
// 		p := &memo[i][hold]
// 		if *p != -1 {
// 			return *p
// 		}

// 		defer func() { *p = res }()

// 		if hold == 1 {
// 			return max(dfs(i-1, 1), dfs(i-1, 0)-prices[i])
// 		}
// 		return max(dfs(i-1, 0), dfs(i-1, 1)+prices[i])

// 	}
// 	return dfs(n-1, 0)
// }

func max(a, b int) int {
	if a > b {
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
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,4,3,1]\n
// @lcpr case=end

*/
