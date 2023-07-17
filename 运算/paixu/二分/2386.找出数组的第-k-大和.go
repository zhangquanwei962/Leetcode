/*
 * @lc app=leetcode.cn id=2386 lang=golang
 * @lcpr version=21908
 *
 * [2386] 找出数组的第 K 大和
 */

// @lc code=start
// sum一定是最大的子序列和，然后全部变为正数减去，减去第k-1
// 小的子序列和，肯定就是第k大子序列和
package main

import (
	"sort"
)

func kSum(nums []int, k int) int64 {
	n := len(nums)
	sum, tot := 0, 0

	for i, x := range nums {
		if x >= 0 {
			sum += x
			tot += x
		} else {
			tot -= x
			nums[i] = -x
		}
	}

	sort.Ints(nums)
	k--

	kthSmallest := sort.Search(tot, func(limit int) bool {
		cnt := 0
		var dfs func(int, int)
		dfs = func(i, s int) {
			if i == n || cnt >= k || s+nums[i] > limit {
				return
			}
			cnt++
			dfs(i+1, s+nums[i]) // select
			dfs(i+1, s)         // not select
		}
		dfs(0, 0)
		return cnt >= k
	})

	return int64(sum - kthSmallest)
}

// 时间复杂度：O(nlogn+klogU)
// O(min(k,n))

// @lc code=end

/*
// @lcpr case=start
// [2,4,-2]\n5\n
// @lcpr case=end

// @lcpr case=start
// [1,-2,3,4,-10,12]\n16\n
// @lcpr case=end

*/
