/*
 * @lc app=leetcode.cn id=1406 lang=golang
 * @lcpr version=21909
 *
 * [1406] 石子游戏 III
 */

// @lc code=start
package main

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + stoneValue[i]
	}

	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int
	dfs = func(i int) (res int) {
		if i == n {
			return
		}

		p := &memo[i]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()

		res = pre[n] - pre[i] - dfs(i+1)
		// 如果还剩余至少两个石头，请拿走
		if i+1 < n {
			res = max(res, pre[n]-pre[i]-dfs(i+2))
		}
		// 如果还剩余至少三个石头，请拿走
		if i+2 < n {
			res = max(res, pre[n]-pre[i]-dfs(i+3))
		}
		return
	}

	Alice := dfs(0)
	Bob := pre[n] - Alice

	if Alice > Bob {
		return "Alice"
	} else if Alice == Bob {
		return "Tie"
	}
	return "Bob"
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
// [1,2,3,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,-9]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,6]\n
// @lcpr case=end

*/
