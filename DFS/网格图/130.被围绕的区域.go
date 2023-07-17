/*
 * @lc app=leetcode.cn id=130 lang=golang
 * @lcpr version=21909
 *
 * [130] 被围绕的区域
 */

// @lc code=start
package main

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	n, m := len(board), len(board[0])
	var dfs func(int, int)
	dfs = func(x, y int) {
		if x < 0 || x >= n || y < 0 || y >= m || board[x][y] != byte('O') {
			return
		}
		board[x][y] = 'A'
		for _, p := range dirs4 {
			x, y = x+p[0], y+p[1]
			dfs(x, y)
		}
	}
	for i := 0; i < n; i++ {
		dfs(i, 0)
		dfs(i, m-1)
	}
	for i := 0; i < m; i++ {
		dfs(0, i)
		dfs(n-1, i)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

var dirs4 = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// @lc code=end

/*
// @lcpr case=start
// [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]\n
// @lcpr case=end

// @lcpr case=start
// [["X"]]\n
// @lcpr case=end

*/
