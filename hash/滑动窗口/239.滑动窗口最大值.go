/*
 * @lc app=leetcode.cn id=239 lang=golang
 * @lcpr version=21909
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
package main

func maxSlidingWindow(nums []int, k int) (res []int) {
	q := []int{}

	for i, x := range nums {
		// 超过区间长度，删除头
		if len(q) > 0 && i-k+1 > q[0] {
			q = q[1:]
		}
		// 尾部小于等于当前x，删去
		for len(q) > 0 && nums[q[len(q)-1]] <= x {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if i >= k-1 {
			res = append(res, nums[q[0]])
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [1,3,-1,-3,5,3,6,7]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

*/
