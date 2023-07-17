/*
 * @lc app=leetcode.cn id=1140 lang=golang
 * @lcpr version=21909
 *
 * [1140] 石子游戏 II
 */

// @lc code=start
package main

func stoneGameII(piles []int) int {
	n := len(piles)
	pre := make([]int, n+1)
	for i := 0; i < n; i++ { // 前缀和
		pre[i+1] = pre[i] + piles[i]
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(i int, m int) (res int) {
		M := m * 2

		if M >= n-i {
			return pre[n] - pre[i]
		}

		p := &memo[i][m]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		for x := 1; x <= M; x++ {
			next_m := max(m, x)
			cur := pre[n] - pre[i] - dfs(i+x, next_m)
			res = max(res, cur)
		}
		return
	}
	return dfs(0, 1)
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
// [2,7,9,4,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5,100]\n
// @lcpr case=end

*/
