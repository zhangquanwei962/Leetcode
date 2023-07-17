/*
 * @lc app=leetcode.cn id=2657 lang=golang
 * @lcpr version=21907
 *
 * [2657] 找到两个数组的前缀公共数组
 */

// @lc code=start
// 集合的观点，
package main

import "math/bits"

func findThePrefixCommonArray(A []int, B []int) []int {
	ans := make([]int, len(A))
	var p, q uint
	for i, x := range A {
		p |= 1 << uint(x)
		q |= 1 << B[i]
		ans[i] = bits.OnesCount(p & q)
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,3,2,4]\n[3,1,2,4]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,1]\n[3,1,2]\n
// @lcpr case=end

*/
