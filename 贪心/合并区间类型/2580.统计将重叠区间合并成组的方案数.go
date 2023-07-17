/*
 * @lc app=leetcode.cn id=2580 lang=golang
 * @lcpr version=21909
 *
 * [2580] 统计将重叠区间合并成组的方案数
 */

// @lc code=start
package main

import "sort"

func countWays(ranges [][]int) int {
	const mod int = 1e9 + 7
	sort.Slice(ranges, func(i, j int) bool { return ranges[i][0] < ranges[j][0] })
	ans, maxR := 2, ranges[0][1]
	for _, p := range ranges[1:] {
		if p[0] > maxR {
			ans = ans * 2 % mod
		}
		maxR = max(maxR, p[1])
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [[6,10],[5,15]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,3],[10,20],[2,5],[4,8]]\n
// @lcpr case=end

*/
