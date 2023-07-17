/*
 * @lc app=leetcode.cn id=329 lang=golang
 * @lcpr version=21909
 *
 * [329] 矩阵中的最长递增路径
 */

// @lc code=start
package main

func longestIncreasingPath(mat [][]int) (res int) {
	m, n := len(mat), len(mat[0])
	memo := make([][]int, m+1)
	dirs4 := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(x, y int) int {
		p := &memo[x][y]
		if *p != -1 {
			return *p
		}

		res := 1

		defer func() { *p = res }()

		for i := 0; i < 4; i++ {
			a, b := x+dirs4[i][0], y+dirs4[i][1]
			if a >= 0 && a < m && b >= 0 && b < n && mat[a][b] > mat[x][y] {
				res = max(res, dfs(a, b)+1)
			}
		}
		return res
	}

	for i := range mat {
		for j := range mat[0] {
			res = max(res, dfs(i, j))
		}
	}
	return
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
// [[9,9,4],[6,6,8],[2,1,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[3,4,5],[3,2,6],[2,2,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1]]\n
// @lcpr case=end

*/
