/*
 * @lc app=leetcode.cn id=2369 lang=golang
 * @lcpr version=21907
 *
 * [2369] 检查数组是否存在有效划分
 */

// @lc code=start
package main

func validPartition(nums []int) bool {
	n := len(nums)
	f := make([]bool, n+1)
	f[0] = true

	for i, x := range nums {
		if i > 0 && f[i-1] && nums[i-1] == x ||
			i > 1 && f[i-2] && (x == nums[i-1] && x == nums[i-2] ||
				x == nums[i-1]+1 && x == nums[i-2]+2) {
			f[i+1] = true
		}
	}
	return f[n]
}

// @lc code=end

/*
// @lcpr case=start
// [4,4,4,5,6]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,2]\n
// @lcpr case=end

*/
