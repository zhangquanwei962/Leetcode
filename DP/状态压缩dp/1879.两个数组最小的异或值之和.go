/*
 * @lc app=leetcode.cn id=1879 lang=golang
 * @lcpr version=21909
 *
 * [1879] 两个数组最小的异或值之和
 */

// @lc code=start
package main

import "math/bits"

func minimumXORSum(x []int, y []int) int {
	m := 1 << len(x)
	dp := make([]int, m)
	for i := range dp {
		dp[i] = 2e9
	}

	dp[0] = 0
	for s, dv := range dp[:m-1] {
		v := x[bits.OnesCount(uint(s))]
		// 将二进制数 t 的低m位全部取反
		for t, lb := s^(m-1), 0; t > 0; t ^= lb {
			lb = t & -t // 计算出二进制数 t 的最低位 1 所代表的数值。
			w := y[bits.TrailingZeros(uint(lb))]
			dp[s|lb] = min(dp[s|lb], v^w+dv)
		}
	}
	return dp[m-1]
}

//通过不断地使用 t ^= lb 来将二进制数 t 的最低位 1
//所代表的数值依次取反，可以依次获取到数组 y 中的每个元素，
//从而枚举所有可能的组合方式，计算最小异或和。
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [1,2]\n[2,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,0,3]\n[5,3,4]\n
// @lcpr case=end

*/
