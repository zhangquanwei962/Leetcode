/*
 * @lc app=leetcode.cn id=11 lang=golang
 * @lcpr version=21909
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
package main

func maxArea(height []int) (ans int) {
	left, right := 0, len(height)-1
	for left < right {
		area := (right - left) * min(height[right], height[left])
		ans = max(ans, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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
// [1,8,6,2,5,4,8,3,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,1]\n
// @lcpr case=end

*/
