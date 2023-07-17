/*
 * @lc app=leetcode.cn id=456 lang=golang
 * @lcpr version=21909
 *
 * [456] 132 模式
 */

// @lc code=start
// **枚举3** 找到比3大的2 如果前缀最小值小于3，就是正确
package main

import (
	"math"
)

func find132pattern(nums []int) bool {
	n := len(nums)
	st := []int{}
	pmin := []int{math.MaxInt32}

	for i := 0; i < n; i++ {
		for len(st) > 0 && nums[st[len(st)-1]] <= nums[i] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 && pmin[st[len(st)-1]] < nums[i] {
			return true
		}
		st = append(st, i)
		pmin = append(pmin, min(pmin[len(pmin)-1], nums[i]))
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [3,1,4,2]\n
// @lcpr case=end

// @lcpr case=start
// [-1,3,2,0]\n
// @lcpr case=end

*/
