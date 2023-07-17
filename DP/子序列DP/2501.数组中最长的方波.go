/*
 * @lc app=leetcode.cn id=2501 lang=golang
 * @lcpr version=21908
 *
 * [2501] 数组中最长的方波
 */

// @lc code=start
package main

func longestSquareStreak(nums []int) (ans int) {
	set := map[int]bool{}

	for _, x := range nums {
		set[x] = true
	}

	dp := map[int]int{}
	var dfs func(int) int
	dfs = func(x int) int {
		if !set[x] {
			return 0
		}

		if v, ok := dp[x]; ok {
			return v
		}
		dp[x] = 1 + dfs(x*x)
		return dp[x]
	}

	for x := range set {
		ans = max(ans, dfs(x))
	}

	if ans == 1 {
		return -1
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [4,3,6,16,8,2]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,5,6,7]\n
// @lcpr case=end

*/
