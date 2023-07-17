/*
 * @lc app=leetcode.cn id=807 lang=golang
 * @lcpr version=21909
 *
 * [807] 保持城市天际线
 */

// @lc code=start
package main

func maxIncreaseKeepingSkyline(grid [][]int) (ans int) {
	n := len(grid)
	rowMax := make([]int, n)
	colMax := make([]int, n)
	for i, row := range grid {
		for j, h := range row {
			rowMax[i] = max(rowMax[i], h)
			colMax[j] = max(colMax[j], h)
		}
	}
	for i, row := range grid {
		for j, h := range row {
			ans += min(rowMax[i], colMax[j]) - h
		}
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
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
// [[3,0,8,4],[2,4,5,7],[9,2,6,3],[0,3,1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,0,0],[0,0,0],[0,0,0]]\n
// @lcpr case=end

*/
