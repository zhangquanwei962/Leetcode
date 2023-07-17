/*
 * @lc app=leetcode.cn id=523 lang=golang
 * @lcpr version=21909
 *
 * [523] 连续的子数组和
 */

// @lc code=start
package main

func checkSubarraySum(nums []int, k int) bool {
	n, remind := len(nums), 0
	if n < 2 {
		return false
	}

	mp := map[int]int{0: -1} //保证前缀和为0时候的下标

	for i, v := range nums {
		remind = (remind + v) % k
		if j, ok := mp[remind]; ok {
			if i-j >= 2 {
				return true
			}
		} else {
			mp[remind] = i
		}
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// [23,2,4,6,7]\n6\n
// @lcpr case=end

// @lcpr case=start
// [23,2,6,4,7]\n6\n
// @lcpr case=end

// @lcpr case=start
// [23,2,6,4,7]\n13\n
// @lcpr case=end

*/
