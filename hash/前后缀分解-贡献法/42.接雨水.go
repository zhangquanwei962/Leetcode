/*
 * @lc app=leetcode.cn id=42 lang=golang
 * @lcpr version=21908
 *
 * [42] 接雨水
 */

// @lc code=start
package main

func trap(height []int) (ans int) {
	n := len(height)
	proMax := make([]int, n)
	proMax[0] = height[0]

	for i := 1; i < n; i++ {
		proMax[i] = max(proMax[i-1], height[i])
	}

	sufMax := make([]int, n)
	sufMax[n-1] = height[n-1]

	for i := n - 2; i >= 0; i-- {
		sufMax[i] = max(sufMax[i+1], height[i])
	}

	for i, h := range height {
		ans += min(sufMax[i], proMax[i]) - h
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
	if a < b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [0,1,0,2,1,0,1,3,2,1,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [4,2,0,3,2,5]\n
// @lcpr case=end

*/
