/*
 * @lc app=leetcode.cn id=面试题 16.19 lang=golang
 * @lcpr version=21909
 *
 * [面试题 16.19] 水域大小
 */

// @lc code=start
package main

import "sort"

var dirs4 = [8][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

func pondSizes(land [][]int) []int {
	m, n := len(land), len(land[0])
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || land[i][j] != 0 {
			return 0
		}
		land[i][j] = -1
		res := 1
		for _, p := range dirs4 {
			x, y := i+p[0], j+p[1]
			res += dfs(x, y)
		}
		return res
	}
	res := []int{}
	for i, r := range land {
		for j, c := range r {
			if c == 0 {
				res = append(res, dfs(i, j))
			}
		}
	}
	sort.Ints(res)
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [[0,2,1,0],[0,1,0,1],[1,1,0,1],[0,1,0,1]]\n
// @lcpr case=end

*/
