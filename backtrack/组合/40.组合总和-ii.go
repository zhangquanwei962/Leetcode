/*
 * @lc app=leetcode.cn id=40 lang=golang
 * @lcpr version=21909
 *
 * [40] 组合总和 II
 */

// @lc code=start
// 只能用一次
package main

import (
	"sort"
)

func combinationSum2(candidates []int, target int) (ans [][]int) {
	n := len(candidates)
	sort.Ints(candidates)
	path := []int{}
	var dfs func(int, int)
	dfs = func(i, sub int) {
		if sub == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for j := i; j >= 0; j-- {
			if sub-candidates[j] < 0 || j < i && candidates[j] == candidates[j+1] {
				continue
			}
			path = append(path, candidates[j])
			dfs(j-1, sub-candidates[j])
			path = path[:len(path)-1]
		}
	}

	dfs(n-1, target)
	return
}

// @lc code=end

/*
// @lcpr case=start
// [10,1,2,7,6,1,5]\n8\n
// @lcpr case=end

// @lcpr case=start
// [2,5,2,1,2]\n5\n
// @lcpr case=end

*/
