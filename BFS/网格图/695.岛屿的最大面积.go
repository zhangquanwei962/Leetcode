/*
 * @lc app=leetcode.cn id=695 lang=golang
 * @lcpr version=21908
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
package main

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				res := bfs(grid, i, j)
				ans = max(ans, res)
			}
		}
	}
	return ans
}

func bfs(grid [][]int, x, y int) int {
	m, n := len(grid), len(grid[0])
	q := []int{x*n + y}
	grid[x][y] = 0
	ans := 1
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, p := range tmp {
			i, j := p/n, p%n
			for k := 0; k < 4; k++ {
				a, b := i+dx[k], j+dy[k]
				if a >= 0 && a < m && b >= 0 && b < n && grid[a][b] == 1 {
					grid[a][b] = 0
					ans++
					q = append(q, a*n+b)
				}
			}
		}

	}
	return ans
}

var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func maxAreaOfIsland(grid [][]int) (ans int) {
// 	m, n := len(grid), len(grid[0])
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if grid[i][j] == 1 {
// 				q := [][2]int{{i, j}}
// 				grid[i][j] = 0
// 				res := 1
// 				for len(q) > 0 {
// 					tmp := q
// 					q = nil
// 					for _, p := range tmp {
// 						for k := 0; k < 4; k++ {
// 							a, b := p[0]+dx[k], p[1]+dy[k]
// 							if a >= 0 && a < m && b >= 0 && b < n && grid[a][b] == 1 {
// 								grid[a][b] = 0
// 								res++
// 								q = append(q, [2]int{a, b})
// 							}
// 						}
// 					}
// 				}
// 				ans = max(ans, res)
// 			}
// 		}
// 	}
// 	return
// }

// var dx = []int{-1, 0, 1, 0}
// var dy = []int{0, 1, 0, -1}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// @lc code=end

/*
// @lcpr case=start
// [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,0,0,0,0,0,0,0]]\n
// @lcpr case=end

*/
