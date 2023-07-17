/*
 * @lc app=leetcode.cn id=1695 lang=golang
 * @lcpr version=21908
 *
 * [1695] 删除子数组的最大得分
 */

// @lc code=start
package main

func maximumUniqueSubarray(nums []int) (ans int) {
	left, s := 0, 0
	cnt := map[int]int{}
	for right, x := range nums {
		s += x
		cnt[x]++

		for len(cnt) != right-left+1 {
			cnt[nums[left]]--
			s -= nums[left]
			if cnt[nums[left]] == 0 {
				delete(cnt, nums[left])
			}
			left++
		}
		ans = max(ans, s)
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [4,2,4,5,6]\n
// @lcpr case=end

// @lcpr case=start
// [5,2,1,2,5,2,1,2,5]\n
// @lcpr case=end

*/
