/*
 * @lc app=leetcode.cn id=2500 lang=golang
 * @lcpr version=21908
 *
 * [2500] 删除每行中的最大值
 */

// @lc code=start
package main

import "sort"

func deleteGreatestValue(grid [][]int) (ans int) {
	for _, row := range grid {
		sort.Ints(row)
	}

	for j := range grid[0] {
		mx := 0
		for _, row := range grid {
			mx = max(mx, row[j])
		}
		ans += mx
	}
	return

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,4],[3,3,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[10]]\n
// @lcpr case=end

*/
