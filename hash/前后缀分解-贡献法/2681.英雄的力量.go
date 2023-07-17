/*
 * @lc app=leetcode.cn id=2681 lang=golang
 * @lcpr version=21909
 *
 * [2681] 英雄的力量
 */

// @lc code=start
// 一个数的贡献是x^3, 枚举x之前的数作为最小值贡献是
// x^3+x^2 * s = x^2(x+s)
// s = 2 * s + x
package main

import "sort"

func sumOfPower(nums []int) (ans int) {
	const mod int = 1e9 + 7
	sort.Ints(nums)
	s := 0
	for _, x := range nums {
		ans = (ans + x*x%mod*(x+s)) % mod // 中间模一次防止溢出
		s = (s*2 + x) % mod
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [2,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1]\n
// @lcpr case=end

*/
