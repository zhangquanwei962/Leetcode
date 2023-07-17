/*
 * @lc app=leetcode.cn id=2439 lang=golang
 * @lcpr version=21909
 *
 * [2439] 最小化数组中的最大值
 */

// @lc code=start
// 最小化最大值 返回right
// true的时候是right
package main

import (
	"sort"
)

func minimizeArrayValue(nums []int) int {
	mx := 0
	for _, x := range nums {
		mx = max(mx, x)
	}

	// check := func(limit int) bool {
	// 	extra := 0
	// 	for i := len(nums) - 1; i > 0; i-- {
	// 		extra = max(nums[i]+extra-limit, 0)
	// 	}
	// 	return nums[0]+extra <= limit
	// }

	// left, right := -1, mx+1
	// binary := func(left, right int) int {
	// 	for left+1 < right {
	// 		mid := left + (right-left)>>1
	// 		if check(mid) {
	// 			right = mid
	// 		} else {
	// 			left = mid
	// 		}
	// 	}
	// 	return right
	// }
	// return binary(left, right)
	// 最小值 是xxxx √√√
	// 如果是寻找数组的数，小于等于一个没有的数，就是边界
	// 从小到大就得>= 从大到小 <=
	// 不是查询数组久随意了
	return sort.Search(mx, func(limit int) bool {
		extra := 0
		for i := len(nums) - 1; i > 0; i-- {
			extra = max(nums[i]+extra-limit, 0)
		}
		return nums[0]+extra <= limit
	})
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [3,7,1,6]\n
// @lcpr case=end

// @lcpr case=start
// [10,1]\n
// @lcpr case=end

*/
