/*
 * @lc app=leetcode.cn id=2461 lang=golang
 * @lcpr version=21909
 *
 * [2461] 长度为 K 子数组中的最大和
 */

// @lc code=start
package main

import (
	"math"
)

func maximumSubarraySum(nums []int, k int) int64 {
	cnt := make(map[int]int)
	s := 0
	left := 0
	ans := 0
	for right, x := range nums {
		s += x
		cnt[x]++
		for right-left+1 > k {
			s -= nums[left]
			cnt[nums[left]]--
			if cnt[nums[left]] == 0 {
				delete(cnt, nums[left])
			}
			left++
		}
		if right-left+1 == k && len(cnt) == k {
			ans = int(math.Max(float64(ans), float64(s)))
		}
	}
	return int64(ans)
}

// @lc code=end

/*
// @lcpr case=start
// [1,5,4,2,9,9,9]\n3\n
// @lcpr case=end

// @lcpr case=start
// [4,4,4]\n3\n
// @lcpr case=end

*/
