/*
 * @lc app=leetcode.cn id=2321 lang=golang
 * @lcpr version=21909
 *
 * [2321] 拼接数组的最大分数
 */

// @lc code=start
package main

func solve(nums1, nums2 []int) int {
	s1, maxSum, s := 0, 0, 0
	for i, x := range nums1 {
		s1 += x
		s = max(s+nums2[i]-x, 0)
		maxSum = max(maxSum, s)
	}
	return s1 + maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	return max(solve(nums1, nums2), solve(nums2, nums1))
}

// @lc code=end

/*
// @lcpr case=start
// [60,60,60]\n[10,90,10]\n
// @lcpr case=end

// @lcpr case=start
// [20,40,20,70,30]\n[50,20,50,40,20]\n
// @lcpr case=end

// @lcpr case=start
// [7,11,13]\n[1,1,1]\n
// @lcpr case=end

*/
