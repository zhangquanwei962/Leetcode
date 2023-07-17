/*
 * @lc app=leetcode.cn id=877 lang=golang
 * @lcpr version=21909
 *
 * [877] 石子游戏
 */

// @lc code=start
package main

func stoneGame(piles []int) bool {
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
	dfs = func(i, j int) (res int) {
		if i > j {
			return
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		if i == j {
			return piles[i]
		}
		// 选i处，x+pre[j+1]-pre[i+1]-dfs
		// 选j处，x+pre[j]-pre[i]-dfs
		res = max(piles[i]+pre[j+1]-pre[i+1]-dfs(i+1,
			j), piles[j]+pre[j]-pre[i]-dfs(i, j-1))
		return
	}
	return dfs(0, n-1) > pre[n]/2
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
// [5,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [3,7,2,3]\n
// @lcpr case=end

*/
