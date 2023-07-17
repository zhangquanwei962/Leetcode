/*
 * @lc app=leetcode.cn id=31 lang=golang
 * @lcpr version=21909
 *
 * [31] 下一个排列
 */

// @lc code=start
// 排列
package main

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	// 从后向前查找第一个a[i]<a[i+1]，较小数
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := n - 1 // [i+1,n)从后向前找到第一个a[i]<a[j],较大数
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	reverse(nums[i+1:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,5]\n
// @lcpr case=end

*/
