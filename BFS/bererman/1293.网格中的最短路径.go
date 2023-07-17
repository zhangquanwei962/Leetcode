/*
 * @lc app=leetcode.cn id=1293 lang=golang
 * @lcpr version=21909
 *
 * [1293] 网格中的最短路径
 */

// @lc code=start
package main

import (
	"math"
)

func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	if k >= m+n-3 {
		return m + n - 2
	}
	// 三维记忆数组， [x][y][cnt1] 分别表示坐标和当前墙的数量，存储的值为当前状态下遍历过的最短路径长度
	// 当坐标和当前墙的数量相等时，继续带入 dfs， 递归的路径是一样的
	// 所以如果某状态下，当前路径长度大于等于 vis 记录的路径长度，继续递归，只会得到更大或相等的结果，所以直接返回

	vis := make([][][]int, m)
	for i := range vis {
		vis[i] = make([][]int, n)
		for j := range vis[i] {
			vis[i][j] = make([]int, k+1)
			for l := range vis[i][j] {
				vis[i][j][l] = math.MaxInt
			}
		}
	}
	var dirs4 = [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	ans := math.MaxInt
	var dfs func(int, int, int, int)
	dfs = func(x, y, cnt, wallCnt int) {
		if grid[x][y] == 1 {
			wallCnt++
		}
		cnt++
		if wallCnt > k {
			return
		}
		if cnt >= ans {
			return
		}

		if x == m-1 && y == n-1 {
			ans = cnt
			return
		}
		if cnt >= vis[x][y][wallCnt] {
			return
		}

		vis[x][y][wallCnt] = cnt
		for i := 0; i < 4; i++ {
			a, b := x+dirs4[i][0], y+dirs4[i][1]
			if a >= 0 && a < m && b >= 0 && b < n {
				dfs(a, b, cnt, wallCnt)
			}
		}
	}
	dfs(0, 0, -1, 0)
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

// func shortestPath(grid [][]int, k int) int {
// 	const inf = math.MaxInt / 2
// 	m, n := len(grid), len(grid[0])
// 	if k >= m+n-3 {
// 		return m + n - 2
// 	}
// 	f := make([]int, m*n)
// 	for i := range f {
// 		f[i] = inf
// 	}

// 	f[0] = 0
// 	for t := 0; t < m*n; t++ {
// 		if f[m*n-1] <= k {
// 			return t
// 		}
// 		g := make([]int, m*n)
// 		for i := range g {
// 			g[i] = inf
// 		}

// 		for i := 0; i < m; i++ {
// 			for j := 0; j < n; j++ {
// 				if i+1 < m {
// 					g[i*n+j+n] = min(g[i*n+j+n], f[i*n+j]+grid[i+1][j])
// 					g[i*n+j] = min(g[i*n+j], f[i*n+j+n]+grid[i][j])
// 				}

// 				if j+1 < n {
// 					g[i*n+j+1] = min(g[i*n+j+1], f[i*n+j]+grid[i][j+1])
// 					g[i*n+j] = min(g[i*n+j], f[i*n+j+1]+grid[i][j])
// 				}
// 			}
// 		}
// 		f = g
// 	}
// 	return -1
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// @lc code=end

/*
// @lcpr case=start
// [[0,0,0],[1,1,0],[0,0,0],[0,1,1],[0,0,0]]\n1\n
// @lcpr case=end

// @lcpr case=start
// [[0,1,1],[1,1,1],[1,0,0]]\n1\n
// @lcpr case=end

*/
