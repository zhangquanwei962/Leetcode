/*
 * @lc app=leetcode.cn id=1493 lang=golang
 * @lcpr version=21908
 *
 * [1493] 删掉一个元素以后全为 1 的最长子数组
 */

// @lc code=start
package main

func longestSubarray(nums []int) (ans int) {
	left, cnt0 := 0, 0
	for right, x := range nums {
		cnt0 += 1 - x
		for cnt0 > 1 {
			cnt0 -= 1 - nums[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans - 1
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
// [1,1,0,1]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,1,1,0,1,1,0,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1]\n
// @lcpr case=end

*/
