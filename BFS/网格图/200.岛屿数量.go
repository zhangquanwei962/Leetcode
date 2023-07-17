/*
 * @lc app=leetcode.cn id=200 lang=golang
 * @lcpr version=21908
 *
 * [200] 岛屿数量
 */

// @lc code=start
package main

func numIslands(grid [][]byte) (ans int) {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				bfs(grid, i, j)
				ans++
			}
		}
	}
	return
}

func bfs(grid [][]byte, x, y int) {
	m, n := len(grid), len(grid[0])
	q := []int{x*n + y}
	grid[x][y] = '0'
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, p := range tmp {
			i, j := p/n, p%n
			for k := 0; k < 4; k++ {
				a, b := i+dx[k], j+dy[k]
				if a >= 0 && a < m && b >= 0 && b < n && grid[a][b] == '1' {
					grid[a][b] = '0'
					q = append(q, a*n+b)
				}
			}
		}

	}
	return
}

var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]\n
// @lcpr case=end

*/
