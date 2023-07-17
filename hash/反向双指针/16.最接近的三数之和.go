/*
 * @lc app=leetcode.cn id=16 lang=golang
 * @lcpr version=21909
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	minn := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	maxn := math.MaxInt
	for i, x := range nums[:n-2] {
		j, k := i+1, n-1
		for j < k {
			s := x + nums[k] + nums[j]
			if s < target {
				j++
			} else if s > target {
				k--
			} else {
				return s
			}

			if maxn > abs(s-target) {
				minn = s
				maxn = abs(s - target)
			}
		}
	}
	return minn
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [-1,2,1,-4]\n1\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0]\n1\n
// @lcpr case=end

*/
