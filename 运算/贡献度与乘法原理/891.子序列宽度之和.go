/*
 * @lc app=leetcode.cn id=891 lang=golang
 * @lcpr version=21909
 *
 * [891] 子序列宽度之和
 */

// @lc code=start
// 与顺序无关，可以排序
package main

import "sort"

const mod int = 1e9 + 7

func sumSubseqWidths(nums []int) (ans int) {
	sort.Ints(nums)
	n := len(nums)
	for i, x := range nums {
		ans += (pow(2, i) - pow(2, n-1-i)) * x // 在题目的数据范围下，这不会溢出
	}
	return (ans%mod + mod) % mod // 注意上面有减法，ans 可能为负数
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n&1 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [2]\n
// @lcpr case=end

*/
