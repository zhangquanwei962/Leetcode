/*
 * @lc app=leetcode.cn id=88 lang=golang
 * @lcpr version=21908
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
package main

func merge(nums1 []int, m int, nums2 []int, n int) {

	for i, j, tail := m-1, n-1, m+n-1; i >= 0 || j >= 0; tail-- {
		var cur int

		if i == -1 {
			cur = nums2[j]
			j--
		} else if j == -1 {
			cur = nums1[i]
			i--
		} else if nums1[i] > nums2[j] {
			cur = nums1[i]
			i--
		} else {
			cur = nums2[j]
			j--
		}
		nums1[tail] = cur
	}
	return

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,0,0,0]\n3\n[2,5,6]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n[]\n0\n
// @lcpr case=end

// @lcpr case=start
// [0]\n0\n[1]\n1\n
// @lcpr case=end

*/
