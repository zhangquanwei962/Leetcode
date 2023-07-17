/*
 * @lc app=leetcode.cn id=2576 lang=golang
 * @lcpr version=21909
 *
 * [2576] 求出最多标记下标
 */

// @lc code=start
// 与序号无关，排序
// 如果可以匹配
// �
// k 对，那么也可以匹配小于
// �
// k 对，去掉一些数对即可做到。

// 如果无法匹配
// �
// k 对，那么也无法匹配大于
// �
// k 对（反证法）。
// **<=**
package main

import "sort"

func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	check := func(mid int) bool {
		for i, x := range nums[:mid] {
			if x*2 > nums[n-mid+i] {
				return false
			}
		}
		return true
	}

	left, right := 0, n/2+1
	for left+1 < right {
		mid := left + (right-left)/2
		if check(mid) {
			left = mid
		} else {
			right = mid
		}
	}
	return left * 2

	// return 2 * sort.Search(n/2, func(k int) bool {
	// 	k++
	// 	for i, x := range nums[:k] {
	// 		if x*2 > nums[n-k+i] {
	// 			return true
	// 		}
	// 	}
	// 	return false
	// })

	//库函数只支持找最小值，那么可以通过 +1
	//转换成找最小的不满足要求的，就得到了最大的满足要求的
}

// @lc code=end

/*
// @lcpr case=start
// [3,5,2,4]\n
// @lcpr case=end

// @lcpr case=start
// [9,2,5,4]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,8]\n
// @lcpr case=end

*/
