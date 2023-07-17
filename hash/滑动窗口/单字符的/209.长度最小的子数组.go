/*
 * @lc app=leetcode.cn id=209 lang=golang
 * @lcpr version=21909
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
package main

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	ans, s, left := n+1, 0, 0
	for right, x := range nums {
		s += x
		for s >= target { // 满足要求
			ans = min(ans, right-left+1)
			s -= nums[left]
			left++
		}

	}
	if ans <= n {
		return ans
	}
	return 0
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
// 7\n[2,3,1,2,4,3]\n
// @lcpr case=end

// @lcpr case=start
// 4\n[1,4,4]\n
// @lcpr case=end

// @lcpr case=start
// 11\n[1,1,1,1,1,1,1,1]\n
// @lcpr case=end

*/
