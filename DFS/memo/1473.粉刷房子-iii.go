/*
 * @lc app=leetcode.cn id=1473 lang=golang
 * @lcpr version=21909
 *
 * [1473] 粉刷房子 III
 */

// @lc code=start
package main

import (
	"math"
)

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	memo := make([][][]int, m+1)
	for i := range memo {
		memo[i] = make([][]int, target+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, n+1)
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}

	var dfs func(int, int, int) int
	dfs = func(i, target, color int) int {
		if i < 0 || target < 0 || i+1 < target { // 不够划分
			if i == -1 && target == 0 {
				return 0
			}
			return math.MaxInt / 2
		}

		p := &memo[i][target][color]
		if *p != -1 {
			return *p
		}
		res := math.MaxInt / 2
		defer func() { *p = res }()
		if houses[i] != 0 {
			if houses[i] == color {
				res = dfs(i-1, target, houses[i])
			} else {
				res = dfs(i-1, target-1, houses[i])
			}
		} else {
			for j := 0; j < n; j++ {
				if j == color-1 {
					res = min(res, dfs(i-1, target, j+1)+cost[i][j])
				} else {
					res = min(res, dfs(i-1, target-1, j+1)+cost[i][j])
				}
			}
		}

		return res
	}

	ans := dfs(m-1, target, 0)
	if ans == math.MaxInt/2 {
		return -1
	}
	return ans

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [0,0,0,0,0]\n[[1,10],[10,1],[10,1],[1,10],[5,1]]\n5\n2\n3\n
// @lcpr case=end

// @lcpr case=start
// [0,2,1,2,0]\n[[1,10],[10,1],[10,1],[1,10],[5,1]]\n5\n2\n3\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0,0,0]\n[[1,10],[10,1],[1,10],[10,1],[1,10]]\n5\n2\n5\n
// @lcpr case=end

// @lcpr case=start
// [3,1,2,3]\n[[1,1,1],[1,1,1],[1,1,1],[1,1,1]]\n4\n3\n3\n
// @lcpr case=end

*/
