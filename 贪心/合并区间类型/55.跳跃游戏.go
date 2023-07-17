/*
 * @lc app=leetcode.cn id=55 lang=golang
 * @lcpr version=21908
 *
 * [55] 跳跃游戏
 */

// @lc code=start
package main

func canJump(nums []int) bool {
	mx := 0
	cur := 0
	for i, x := range nums {
		mx = max(mx, x+i)
		if i == cur {
			if i == mx && i != len(nums)-1 {
				return false
			}
		}
		cur = mx
	}
	return true
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
// [2,3,1,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,1,0,4]\n
// @lcpr case=end

*/
