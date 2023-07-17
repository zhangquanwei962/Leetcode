/*
 * @lc app=leetcode.cn id=69 lang=golang
 * @lcpr version=21909
 *
 * [69] x 的平方根
 */

// @lc code=start
// 目标是<=就找<=,可以<= >互换，都是找到第一个最后一个小于等于得
package main

func mySqrt(x int) int {
	binary := func(left, right int) int {
		for left+1 < right {
			mid := left + (right-left)>>1
			if mid*mid <= x {
				left = mid
			} else {
				right = mid
			}
		}
		return left
	}
	return binary(-1, x+1)
}

// @lc code=end

/*
// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 8\n
// @lcpr case=end

*/
