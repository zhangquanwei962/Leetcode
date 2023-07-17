/*
 * @lc app=leetcode.cn id=47 lang=golang
 * @lcpr version=21909
 *
 * [47] 全排列 II
 */

// @lc code=start
package main

import "sort"

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}
	path := []int{}
	onPath := make([]bool, n)
	sort.Ints(nums)

	var dfs func(i int, idx int)
	dfs = func(i int, idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for j := 0; j < n; j++ {
			if onPath[j] || (j > 0 && nums[j] == nums[j-1] && onPath[j-1]) {
				continue
			}
			path = append(path, nums[j])
			onPath[j] = true
			dfs(j+1, idx+1)
			onPath[j] = false
			path = path[:len(path)-1]
		}
	}

	dfs(0, 0)
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

*/
