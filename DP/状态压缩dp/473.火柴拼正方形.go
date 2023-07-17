/*
 * @lc app=leetcode.cn id=473 lang=golang
 * @lcpr version=21909
 *
 * [473] 火柴拼正方形
 */

// @lc code=start
package main

import "sort"

func makesquare(matchsticks []int) bool {
	s := 0
	for _, x := range matchsticks {
		s += x
	}

	if s%4 != 0 {
		return false
	}
	target := s / 4
	sort.Ints(matchsticks)
	n := len(matchsticks)
	if target < matchsticks[n-1] {
		return false
	}
	mask := 1<<n - 1
	f := make([]int, 1<<n)
	var dfs func(int, int) bool
	dfs = func(state, t int) bool {
		if state == mask {
			return true
		}
		if f[state] != 0 {
			return f[state] == 1
		}

		for i, v := range matchsticks {
			if (state >> i & 1) == 1 {
				continue
			}

			if t+v > target {
				break
			}

			// 取余，刚好可以看是不是相等
			if dfs(state|1<<i, (t+v)%target) {
				f[state] = 1
				return true
			}
		}
		f[state] = -1
		return false
	}
	return dfs(0, 0)
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,2,2,2]\n
// @lcpr case=end

// @lcpr case=start
// [3,3,3,3,4]\n
// @lcpr case=end

*/
