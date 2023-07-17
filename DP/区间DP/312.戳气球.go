/*
 * @lc app=leetcode.cn id=312 lang=golang
 * @lcpr version=21909
 *
 * [312] 戳气球
 */

// @lc code=start
package main

import "math"

func maxCoins(nums []int) int {
	n := len(nums)
	numsNew := append([]int{1}, append(nums, 1)...)
	memo := make([][]int, n+2)
	for i := range memo {
		memo[i] = make([]int, n+2)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i int, j int) int {
		if i+1 == j {
			return 0
		}

		p := &memo[i][j]
		if *p != -1 {
			return *p
		}

		ans := math.MinInt32
		for k := i + 1; k < j; k++ {
			left := dfs(i, k)
			right := dfs(k, j)
			temp := numsNew[i]*numsNew[k]*numsNew[j] + left + right
			if temp > ans {
				ans = temp
			}
		}
		defer func() { *p = ans }()
		return ans
	}
	return dfs(0, n+1)
}

// @lc code=end

/*
// @lcpr case=start
// [3,1,5,8]\n
// @lcpr case=end

// @lcpr case=start
// [1,5]\n
// @lcpr case=end

*/
