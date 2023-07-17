/*
 * @lc app=leetcode.cn id=1401 lang=golang
 * @lcpr version=21909
 *
 * [1401] 圆和矩形是否有重叠
 */

// @lc code=start
package main

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	dist := 0
	if xCenter < x1 || xCenter > x2 {
		dist += min((x1-xCenter)*(x1-xCenter), (x2-xCenter)*(x2-xCenter))
	}
	if yCenter < y1 || yCenter > y2 {
		dist += min((y1-yCenter)*(y1-yCenter), (y2-yCenter)*(y2-yCenter))
	}
	return dist <= radius*radius
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
// 1\n0\n0\n1\n-1\n3\n1\n
// @lcpr case=end

// @lcpr case=start
// 1\n1\n1\n1\n-3\n2\n-1\n
// @lcpr case=end

// @lcpr case=start
// 1\n0\n0\n-1\n0\n0\n1\n
// @lcpr case=end

*/
