/*
 * @lc app=leetcode.cn id=1872 lang=golang
 * @lcpr version=21909
 *
 * [1872] 石子游戏 VIII
 */

// @lc code=start
// 前缀和不变，倒序DP
// f[i]直接定义为最大差
// Bob-Alice = -(Alice-Bob)最大
package main

func stoneGameVIII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}

	f := sum
	for i := len(stones) - 1; i > 1; i-- {
		sum -= stones[i]
		f = max(f, sum-f)
	}
	return f
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [-1,2,-3,4,-5]\n
// @lcpr case=end

// @lcpr case=start
// [7,-6,5,10,5,-2,-6]\n
// @lcpr case=end

// @lcpr case=start
// [-10,-12]\n
// @lcpr case=end

*/
