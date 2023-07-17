/*
 * @lc app=leetcode.cn id=1510 lang=golang
 * @lcpr version=21909
 *
 * [1510] 石子游戏 IV
 */

// @lc code=start
// 记忆化搜索
package main

import (
	"math"
)

func winnerSquareGame(n int) bool {
	memo := map[int]bool{}

	var dfs func(int) bool
	dfs = func(i int) bool {
		if i < 0 {
			return false
		}
		if p, ok := memo[i]; ok {
			return p
		}
		// 能使对方输就是自己赢，不能使对方输就是自己输
		for j := int(math.Sqrt(float64(i))); j >= 1; j-- {
			flag := dfs(i - j*j)
			if !flag {
				memo[i] = true
				return true
			}
		}
		memo[i] = false
		return false
	}
	return dfs(n)
}

// @lc code=end

/*
// @lcpr case=start
// 1\n
// @lcpr case=end

// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 7\n
// @lcpr case=end

// @lcpr case=start
// 17\n
// @lcpr case=end

*/
