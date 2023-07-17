/*
 * @lc app=leetcode.cn id=63 lang=golang
 * @lcpr version=21907
 *
 * [63] 不同路径 II
 */

// @lc code=start
// O(mn) O(mn)
package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[m-1][n-1] == 1 || obstacleGrid[0][0] == 1 {
		return 0
	}

	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没被计算过
		}
	}
	var d = [][2]int{{-1, 0}, {0, -1}}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		// if i < 0 || j < 0 {
		// 	return
		// }
		if j == 0 && i == 0 {
			// memo[i][j] = 1
			return 1
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		// if i-1 >= 0 && obstacleGrid[i-1][j] == 0 {
		// 	res += dfs(i-1, j)
		// }
		// if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
		// 	res += dfs(i, j-1)
		// }
		for _, p := range d {
			x, y := i+p[0], j+p[1]
			if x >= 0 && x < m && y >= 0 && y < n && obstacleGrid[x][y] == 0 {
				res += dfs(x, y)
			}
		}
		return res
	}
	return dfs(m-1, n-1)
}

// @lc code=end

/*
// @lcpr case=start
// [[0,0,0],[0,1,0],[0,0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1],[0,0]]\n
// @lcpr case=end

*/
