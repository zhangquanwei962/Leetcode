/*
 * @lc app=leetcode.cn id=1254 lang=golang
 * @lcpr version=21909
 *
 * [1254] 统计封闭岛屿的数目
 */

// @lc code=start
package main

func closedIsland(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	if m < 3 || n < 3 {
		return
	}
	var closed bool
	var dfs func(int, int)
	dfs = func(x, y int) {
		if x == 0 || x == m-1 || y == 0 || y == n-1 {
			if grid[x][y] == 0 { // 到达边界
				closed = false
			}
			return
		}
		if grid[x][y] != 0 {
			return
		}
		grid[x][y] = 1 // 标记 (x,y) 被访问，避免重复访问
		for _, p := range dirs4 {
			x, y := p[0]+x, p[1]+y
			dfs(x, y)
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 { // 从没有访问过的 0 出发
				closed = true
				dfs(i, j)
				if closed {
					ans++
				}
			}
		}
	}
	return
}

var dirs4 = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

// @lc code=end

/*
// @lcpr case=start
// [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,1,1,1,1,1,1],[1,0,0,0,0,0,1],[1,0,1,1,1,0,1],[1,0,1,0,1,0,1],[1,0,1,1,1,0,1],[1,0,0,0,0,0,1],[1,1,1,1,1,1,1]]\n
// @lcpr case=end

*/
