/*
 * @lc app=leetcode.cn id=611 lang=golang
 * @lcpr version=21909
 *
 * [611] 有效三角形的个数
 */

// @lc code=start
package main

import "sort"

func triangleNumber(nums []int) (ans int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		k := i
		for j := i + 1; j < n; j++ { // 增加时候也增加，找到最大满足nums[k]<nums[i]+nums[j]
			for k+1 < n && nums[k+1] < nums[i]+nums[j] {
				k++
			}
			ans += max(k-j, 0)
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

// func triangleNumber(nums []int) (ans int) {
// 	sort.Ints(nums)
// 	n := len(nums)

// 	for i := n - 1; i > 1; i-- {
// 		j, k := 0, i-1
// 		for j < k {
// 			if nums[j]+nums[k] > nums[i] {
// 				ans += k - j
// 				k--
// 			} else {
// 				j++
// 			}
// 		}
// 	}
// 	return
// }

// @lc code=end

/*
// @lcpr case=start
// [2,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [4,2,3,4]\n
// @lcpr case=end

*/
