/*
 * @lc app=leetcode.cn id=238 lang=golang
 * @lcpr version=21908
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	// left, right := 1, 1
	pre := make([]int, n)
	pre[0] = 1
	for i := 1; i < n; i++ {
		pre[i] = nums[i-1] * pre[i-1]
	}

	suf := make([]int, n)
	suf[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		suf[i] = nums[i+1] * suf[i+1]
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = pre[i] * suf[i]
	}

	return res

	// res := make([]int, n)
	// for i := range res {
	// 	res[i] = 1
	// }
	// for i := 0; i < n; i++ {
	// 	res[i] *= left
	// 	left *= nums[i]

	// 	res[n-i-1] *= right
	// 	right *= nums[n-i-1]
	// }

	// return res

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [-1,1,0,-3,3]\n
// @lcpr case=end

*/
