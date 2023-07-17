/*
 * @lc app=leetcode.cn id=1928 lang=golang
 * @lcpr version=21909
 *
 * [1928] 规定时间内到达终点的最小花费
 */

// @lc code=start
// O((m+n)maxtime) O(n*maxtime)
package main

import "math"

func minCost(maxTime int, edges [][]int, passingFees []int) int {
	const inf = math.MaxInt / 2
	f := make([][]int, maxTime+1)
	n := len(passingFees)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = inf
		}
	}
	f[0][0] = passingFees[0]
	ans := inf
	for t := 1; t <= maxTime; t++ {
		for _, e := range edges {
			i, j, cost := e[0], e[1], e[2]
			if cost <= t {
				f[t][i] = min(f[t][i], f[t-cost][j]+passingFees[i])
				f[t][j] = min(f[t][j], f[t-cost][i]+passingFees[j])
			}
		}
		ans = min(f[t][n-1], ans)
	}

	if ans == inf {
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
// 30\n[[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]]\n[5,1,2,20,20,3]\n
// @lcpr case=end

// @lcpr case=start
// 29\n[[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]]\n[5,1,2,20,20,3]\n
// @lcpr case=end

// @lcpr case=start
// 25\n[[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],[4,5,15]]\n[5,1,2,20,20,3]\n
// @lcpr case=end

*/
