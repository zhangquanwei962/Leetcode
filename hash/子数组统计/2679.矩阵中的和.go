/*
 * @lc app=leetcode.cn id=2679 lang=golang
 * @lcpr version=21909
 *
 * [2679] 矩阵中的和
 */

// @lc code=start
package main

import "sort"

func matrixSum(nums [][]int) (ans int) {
	for _, row := range nums {
		sort.Ints(row)
	}

	for j := range nums[0] {
		mx := 0
		for _, row := range nums {
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
// [[7,2,1],[6,4,2],[6,5,3],[3,2,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1]]\n
// @lcpr case=end

*/
