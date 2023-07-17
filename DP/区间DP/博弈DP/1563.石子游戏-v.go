/*
 * @lc app=leetcode.cn id=1563 lang=golang
 * @lcpr version=21909
 *
 * [1563] 石子游戏 V
 */

// @lc code=start
package main

func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + stoneValue[i]
	}

	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i >= j {
			return
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()

		sum := pre[j+1] - pre[i]
		sumL := 0
		for k := i; k < j; k++ {
			sumL += stoneValue[k]
			sumR := sum - sumL
			if sumR > sumL {
				res = max(res, sumL+dfs(i, k))
			} else if sumR < sumL {
				res = max(res, sumR+dfs(k+1, j))
			} else {
				res = max(res, max(sumL+dfs(i, k), sumR+dfs(k+1, j)))
			}
		}
		return
	}
	return dfs(0, n-1)
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
// [6,2,3,4,5,5]\n
// @lcpr case=end

// @lcpr case=start
// [7,7,7,7,7,7,7]\n
// @lcpr case=end

// @lcpr case=start
// [4]\n
// @lcpr case=end

*/
