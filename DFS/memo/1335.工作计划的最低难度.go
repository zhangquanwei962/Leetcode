/*
 * @lc app=leetcode.cn id=1335 lang=golang
 * @lcpr version=21907
 *
 * [1335] 工作计划的最低难度
 */

// @lc code=start
package main

import "math"

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}
	memo := make([][]int, d)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}

		defer func() { *p = res }()

		if i == 0 {
			for _, x := range jobDifficulty[:j+1] {
				res = max(res, x)
			}
			return
		}

		res = math.MaxInt
		mx := 0
		for k := j; k >= i; k-- {
			mx = max(mx, jobDifficulty[k])
			res = min(res, dfs(i-1, k-1)+mx)
		}
		return
	}
	return dfs(d-1, n-1)
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// @lc code=end

/*
// @lcpr case=start
// [6,5,4,3,2,1]\n2\n
// @lcpr case=end

// @lcpr case=start
// [9,9,9]\n4\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [7,1,7,1,7,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [11,111,22,222,33,333,44,444]\n6\n
// @lcpr case=end

*/
