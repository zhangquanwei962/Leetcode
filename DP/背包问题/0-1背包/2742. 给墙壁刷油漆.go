/*
 * @lc app=leetcode.cn id=2742 lang=golang
 * @lcpr version=21909
 *
 * [2742] 给墙壁刷油漆
 */

// @lc code=start
// 至少型 至少体积为n
// 如果 j<=0 不需要选任何物体，返回0 如果i<0 返回无穷ing大
package main

import (
	"fmt"
	"math"
)

func paintWalls(cost, time []int) int {
	n := len(cost)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2 // 防止加法溢出
	}
	for i, c := range cost {
		t := time[i] + 1 // 注意这里加一了
		for j := n; j > 0; j-- {
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

func main() {
	cost := []int{1, 2, 3, 2}
	time := []int{1, 2, 3, 2}
	fmt.Println(paintWalls(cost, time))
}

// @lc code=end
