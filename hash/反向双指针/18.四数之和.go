/*
 * @lc app=leetcode.cn id=18 lang=golang
 * @lcpr version=21909
 *
 * [18] 四数之和
 */

// @lc code=start
package main

import "sort"

func fourSum(nums []int, target int) (ans [][]int) {
	if len(nums) < 4 {
		return
	}

	n := len(nums)
	sort.Ints(nums)

	for i, x := range nums[:n-3] {
		if i > 0 && x == nums[i-1] {
			continue
		}
		if x+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if x+nums[len(nums)-1]+nums[len(nums)-2]+nums[len(nums)-3] < target {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			y := nums[j]
			if j > i+1 && y == nums[j-1] {
				continue
			}
			if x+y+nums[j+1]+nums[j+2] > target {
				break
			}
			if x+y+nums[len(nums)-1]+nums[len(nums)-2] < target {
				continue
			}
			l, k := j+1, n-1
			for l < k {
				s := x + y + nums[l] + nums[k]
				if s < target {
					l++
				} else if s > target {
					k--
				} else {
					ans = append(ans, []int{x, y, nums[l], nums[k]})
					l++
					for l < k && nums[l] == nums[l-1] {
						l++
					}
					k--
					for l < k && nums[k] == nums[k+1] {
						k--
					}
				}
			}
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [1,0,-1,0,-2,2]\n0\n
// @lcpr case=end

// @lcpr case=start
// [2,2,2,2,2]\n8\n
// @lcpr case=end

*/
