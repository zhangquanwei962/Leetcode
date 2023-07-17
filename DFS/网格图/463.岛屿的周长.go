/*
 * @lc app=leetcode.cn id=463 lang=golang
 * @lcpr version=21908
 *
 * [463] 岛屿的周长
 */

// @lc code=start
package main

type pair struct{ x, y int }

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func islandPerimeter(grid [][]int) (ans int) {
	n, m := len(grid), len(grid[0])
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
			// ans++
			return 1
		}
		if grid[x][y] == 2 {
			return 0
		}
		grid[x][y] = 2
		res := 0 // 初始周长是0
		for _, d := range dir4 {
			res += dfs(x+d.x, y+d.y)
		}
		return res
	}
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				ans += dfs(i, j)
			}
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,0]]\n
// @lcpr case=end

*/
