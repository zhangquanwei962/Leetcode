/*
 * @lc app=leetcode.cn id=2708 lang=golang
 * @lcpr version=21909
 *
 * [2708] 一个小组的最大实力值
 */

// @lc code=start
package main

import "math"

func maxStrength(nums []int) int64 {
	max, min, ans := int64(nums[0]), int64(nums[0]), nums[0]
	for _, v := range nums[1:] {
		tmp, v := max, int64(v)
		max = int64(math.Max(float64(v), math.Max(float64(max), math.Max(float64(max*v), float64(min*v)))))
		min = int64(math.Min(float64(v), math.Min(float64(min), math.Min(float64(tmp*v), float64(min*v)))))

	}
	return int64(math.Max(float64(ans), float64(max)))
}

// @lc code=end

/*
// @lcpr case=start
// [3,-1,-5,2,5,-9]\n
// @lcpr case=end

// @lcpr case=start
// [-4,-5,-4]\n
// @lcpr case=end

*/
