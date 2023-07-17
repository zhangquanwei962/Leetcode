/*
 * @lc app=leetcode.cn id=45 lang=golang
 * @lcpr version=21908
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
package main

func jump(nums []int) (ans int) {
	cur_right, max_right := 0, 0
	for i, x := range nums[:len(nums)-1] {
		max_right = max(max_right, i+x)
		if i == cur_right {
			ans++
			cur_right = max_right
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
// [2,3,1,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,0,1,4]\n
// @lcpr case=end

*/
