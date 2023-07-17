/*
 * @lc app=leetcode.cn id=223 lang=golang
 * @lcpr version=21909
 *
 * [223] 矩形面积
 */

// @lc code=start
package main

func computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 int) int {
	area1 := (ax2 - ax1) * (ay2 - ay1)
	area2 := (bx2 - bx1) * (by2 - by1)
	overlapWidth := min(ax2, bx2) - max(ax1, bx1)
	overlapHeight := min(ay2, by2) - max(ay1, by1)
	overlapArea := max(overlapWidth, 0) * max(overlapHeight, 0)
	return area1 + area2 - overlapArea
}

func min(a, b int) int {
	if a > b {
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

// @lc code=end

/*
// @lcpr case=start
// -3\n0\n3\n4\n0\n-1\n9\n2\n
// @lcpr case=end

// @lcpr case=start
// -2\n-2\n2\n2\n-2\n-2\n2\n2\n
// @lcpr case=end

*/
