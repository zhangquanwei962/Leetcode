/*
 * @lc app=leetcode.cn id=377 lang=golang
 * @lcpr version=21909
 *
 * [377] 组合总和 Ⅳ
 */

// @lc code=start
// 如果组合问题需考虑元素之间的顺序，需将target放在外循环，将nums放在内循环。
package main

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n4\n
// @lcpr case=end

// @lcpr case=start
// [9]\n3\n
// @lcpr case=end

*/
