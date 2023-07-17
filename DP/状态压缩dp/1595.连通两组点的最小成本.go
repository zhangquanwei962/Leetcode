/*
 * @lc app=leetcode.cn id=1595 lang=golang
 * @lcpr version=21909
 *
 * [1595] 连通两组点的最小成本
 */

// @lc code=start
package main

import "math"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func connectTwoGroups(cost [][]int) int {
	m := len(cost[0])
	f := make([]int, 1<<m)
	for j := 0; j < m; j++ {
		mn := math.MaxInt
		for _, c := range cost {
			mn = min(mn, c[j])
		}
		bit := 1 << j
		for mask := 0; mask < bit; mask++ {
			f[bit|mask] = f[mask] + mn
		}
	}
	for _, row := range cost {
		for j := 1<<m - 1; j >= 0; j-- {
			res := math.MaxInt
			for k, c := range row { // 第一组的点 i 与第二组的点 k
				res = min(res, f[j&^(1<<k)]+c)
			}
			f[j] = res
		}
	}
	return f[1<<m-1]
}

// func connectTwoGroups(cost [][]int) int {
// 	n, m := len(cost), len(cost[0])
// 	minCost := make([]int, m)
// 	for j := 0; j < m; j++ {
// 		minCost[j] = math.MaxInt
// 		for _, c := range cost {
// 			minCost[j] = min(minCost[j], c[j])
// 		}
// 	}

// 	f := make([][]int, n+1)
// 	for i := range f {
// 		f[i] = make([]int, 1<<m)
// 	}

// 	for j := 0; j < 1<<m; j++ {
// 		for k, c := range minCost {
// 			if j>>k&1 == 1 { // 第二组的点 k 未连接
// 				f[0][j] += c // 去第一组找个成本最小的点连接
// 			}
// 		}
// 	}

// 	for i, row := range cost {
// 		for j := 0; j < 1<<m; j++ {
// 			res := math.MaxInt
// 			for k, c := range row { // 第一组的点 i 与第二组的点 k
// 				res = min(res, f[i][j&^(1<<k)]+c)
// 			}
// 			f[i+1][j] = res
// 		}
// 	}
// 	return f[n][1<<m-1]

// }

/*
* 记忆化搜索方式 状态压缩dp
 */
// func connectTwoGroups(cost [][]int) int {
// 	n, m := len(cost), len(cost[0])
// 	minCost := make([]int, m)
// 	// 预处理第二组点与第一组点相连的最小cost
// 	for j := 0; j < m; j++ {
// 		minCost[j] = math.MaxInt
// 		for _, c := range cost {
// 			minCost[j] = min(minCost[j], c[j])
// 		}
// 	}

// 	memo := make([][]int, n)
// 	for i := range memo {
// 		memo[i] = make([]int, 1<<m)
// 		for j := range memo[i] {
// 			memo[i][j] = -1
// 		}
// 	}

// 	var dfs func(int, int) int
// 	dfs = func(i, j int) (res int) {
// 		if i < 0 {
// 			for k, c := range minCost {
// 				if j>>k&1 == 1 {
// 					res += c
// 				}
// 			}
// 			return
// 		}

// 		p := &memo[i][j]
// 		if *p != -1 {
// 			return *p
// 		}
// 		defer func() { *p = res }()

// 		res = math.MaxInt
// 		for k, c := range cost[i] {
// 			res = min(res, dfs(i-1, j&^(1<<k))+c)
// 		}
// 		return
// 	}

// 	return dfs(n-1, 1<<m-1)

// }

// @lc code=end

/*
// @lcpr case=start
// [[15, 96], [36, 2]]\n
// @lcpr case=end

// @lcpr case=start
// [[1, 3, 5], [4, 1, 1], [1, 5, 3]]\n
// @lcpr case=end

// @lcpr case=start
// [[2, 5, 1], [3, 4, 7], [8, 1, 2], [6, 2, 4], [3, 8, 8]]\n
// @lcpr case=end

*/
