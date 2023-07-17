/*
 * @lc app=leetcode.cn id=1162 lang=golang
 * @lcpr version=21908
 *
 * [1162] 地图分析
 */

// @lc code=start
package main

func maxDistance(mat [][]int) (ans int) {
	m, n := len(mat), len(mat[0])
	dis := make([][]int, m)

	vis := make([][]bool, m)
	for i := range dis {
		dis[i] = make([]int, n)
		vis[i] = make([]bool, n)
	}
	q := []pair{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				q = append(q, pair{i, j})
				vis[i][j] = true
			}
		}
	}

	if len(q) == 0 || len(q) == m*n {
		return -1
	}

	for len(q) > 0 {
		x, y := q[0].x, q[0].y
		q = q[1:]
		for _, d := range dirs4 {
			a, b := x+d[0], y+d[1]
			if a >= 0 && a < m && b >= 0 && b < n && !vis[a][b] {
				dis[a][b] = dis[x][y] + 1
				q = append(q, pair{a, b})
				vis[a][b] = true
			}
		}
	}

	for i := range mat {
		for j := range mat[0] {
			ans = max(ans, dis[i][j])
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

type pair struct{ x, y int }

var dirs4 = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// @lc code=end

/*
// @lcpr case=start
// [[1,0,1],[0,0,0],[1,0,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,0,0],[0,0,0],[0,0,0]]\n
// @lcpr case=end

*/
