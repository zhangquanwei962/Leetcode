/*
 * @lc app=leetcode.cn id=787 lang=golang
 * @lcpr version=21909
 *
 * [787] K 站中转内最便宜的航班
 */

// @lc code=start
// O((m+n)K) O(nk)
package main

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	const inf = 10000*101 + 1
	f := make([]int, n)
	for i := range f {
		f[i] = inf
	}
	f[src] = 0
	ans := inf
	for t := 1; t <= k+1; t++ {
		g := make([]int, n)
		for i := range g {
			g[i] = inf
		}
		for _, flight := range flights {
			j, i, cost := flight[0], flight[1], flight[2]
			g[i] = min(g[i], f[j]+cost)
		}
		f = g
		ans = min(ans, f[dst])
	}
	if ans == inf {
		ans = -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
// 	const inf = 10000*101 + 1
// 	f := make([][]int, k+2)
// 	for i := range f {
// 		f[i] = make([]int, n)
// 		for j := range f[i] {
// 			f[i][j] = inf
// 		}
// 	}
// 	f[0][src] = 0
// 	for t := 1; t <= k+1; t++ {
// 		for _, flight := range flights {
// 			j, i, cost := flight[0], flight[1], flight[2]
// 			f[t][i] = min(f[t][i], f[t-1][j]+cost)
// 		}
// 	}
// 	ans := inf
// 	for t := 1; t <= k+1; t++ {
// 		ans = min(ans, f[t][dst])
// 	}
// 	if ans == inf {
// 		ans = -1
// 	}
// 	return ans
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// @lc code=end

/*
// @lcpr case=start
// 3\n[[0,1,100],[1,2,100],[0,2\n0\n2\n1\n
// @lcpr case=end

// @lcpr case=start
// 3\n[[0,1,100],[1,2,100],[0,2\n0\n2\n0\n
// @lcpr case=end

*/
