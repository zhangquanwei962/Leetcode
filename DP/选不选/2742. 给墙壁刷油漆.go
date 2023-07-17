/*
 * @lc app=leetcode.cn id=2742 lang=golang
 * @lcpr version=21909
 *
 * [2742] 给墙壁刷油漆
 */

// @lc code=start
package main

import (
	"fmt"
	"math"
)

func paintWalls(cost, time []int) int {
	n := len(cost)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n*2+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if j-n > i { // 注意 j 加上了偏移量 n
			return 0 // 剩余的墙都可以免费刷
		}
		if i < 0 {
			return math.MaxInt / 2 // 防止加法溢出
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		*p = min(dfs(i-1, j+time[i])+cost[i], dfs(i-1, j-1))
		return *p
	}
	return dfs(n-1, n) // 加上偏移量 n，防止负数
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	cost := []int{1, 2, 3, 2}
	time := []int{1, 2, 3, 2}
	fmt.Println(paintWalls(cost, time))
}

// @lc code=end
