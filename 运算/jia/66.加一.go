/*
 * @lc app=leetcode.cn id=66 lang=golang
 * @lcpr version=21909
 *
 * [66] 加一
 */

// @lc code=start
package main

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}

	// 全9
	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [4,3,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/
