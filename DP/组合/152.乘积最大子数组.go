/*
 * @lc app=leetcode.cn id=152 lang=golang
 * @lcpr version=21909
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
package main

func maxProduct(a []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	mi, mx, ans := a[0], a[0], a[0]
	for _, v := range a[1:] {
		mi, mx = min(v, min(v*mi, v*mx)), max(v, max(v*mi, v*mx))
		ans = max(ans, mx)
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,-2,4]\n
// @lcpr case=end

// @lcpr case=start
// [-2,0,-1]\n
// @lcpr case=end

*/
