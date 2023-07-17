/*
 * @lc app=leetcode.cn id=42 lang=golang
 * @lcpr version=21909
 *
 * [42] 接雨水
 */

// @lc code=start
// O(n) O(1)
package main

func trap(height []int) (ans int) {
	left, right, preMax, sufMax := 0, len(height)-1, 0, 0
	for left <= right {
		preMax = max(preMax, height[left])
		sufMax = max(sufMax, height[right])

		if preMax < sufMax {
			ans += preMax - height[left]
			left++
		} else {
			ans += sufMax - height[right]
			right--
		}
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
// [0,1,0,2,1,0,1,3,2,1,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [4,2,0,3,2,5]\n
// @lcpr case=end

*/
