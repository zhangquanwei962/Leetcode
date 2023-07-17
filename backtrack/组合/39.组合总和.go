/*
 * @lc app=leetcode.cn id=39 lang=golang
 * @lcpr version=21909
 *
 * [39] 组合总和
 */

// @lc code=start
// 可以重复用
package main

func combinationSum(candidates []int, target int) (ans [][]int) {
	n := len(candidates)
	path := []int{}
	var dfs func(int, int)
	dfs = func(i, sub int) {
		if sub == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		for j := i; j >= 0; j-- {
			if sub-candidates[j] < 0 {
				continue
			}
			path = append(path, candidates[j])
			dfs(j, sub-candidates[j])
			path = path[:len(path)-1]
		}
	}

	dfs(n-1, target)
	return
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,6,7]\n7\n
// @lcpr case=end

// @lcpr case=start
// [2,3,5]\n8\n
// @lcpr case=end

// @lcpr case=start
// [2]\n1\n
// @lcpr case=end

*/
