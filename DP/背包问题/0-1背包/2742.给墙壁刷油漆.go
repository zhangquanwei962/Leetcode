/*
 * @lc app=leetcode.cn id=2742 lang=golang
 * @lcpr version=21909
 *
 * [2742] 给墙壁刷油漆
 */

// @lc code=start
package main

import (
	"math"
)

func paintWalls(cost, time []int) int {
	n := len(cost)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2 // 防止加法溢出
	}
	for i, c := range cost {
		t := time[i] + 1         // 注意这里加一了
		for j := n; j > 0; j-- { // 0不用算了
			f[j] = min(f[j], f[max(j-t, 0)]+c) // 防止体积为负
		}
	}
	return f[n]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,2]\n[1,2,3,2]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,4,2]\n[1,1,1,1]\n
// @lcpr case=end

*/
