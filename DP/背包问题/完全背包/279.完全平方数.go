/*
 * @lc app=leetcode.cn id=279 lang=golang
 * @lcpr version=21908
 *
 * [279] 完全平方数
 */

// @lc code=start
package main

import "math"

func numSquares(n int) int {
	f := make([]int, n+1)
	for i := range f {
		f[i] = math.MaxInt
	}
	f[0] = 0

	for i := 1; i <= n; i++ {
		x := i * i
		for j := x; j <= n; j++ {
			f[j] = min(f[j], f[j-x]+1)
		}
	}
	return f[n]
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
// 12\n
// @lcpr case=end

// @lcpr case=start
// 13\n
// @lcpr case=end

*/
