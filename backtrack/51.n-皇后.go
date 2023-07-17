/*
 * @lc app=leetcode.cn id=51 lang=golang
 * @lcpr version=21907
 *
 * [51] N 皇后
 */

// @lc code=start
package main

import "strings"

func solveNQueens(n int) (ans [][]string) {
	col := make([]int, n)
	onPath := make([]bool, n)
	dig1 := make([]bool, n*2-1)
	dig2 := make([]bool, n*2-1)

	var dfs func(int)
	dfs = func(r int) {
		if r == n {
			board := make([]string, n)
			for i, c := range col {
				board[i] = strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-1-c)
			}
			ans = append(ans, board)
			return
		}

		for c, on := range onPath {
			rc := r - c + n - 1
			if !on && !dig1[r+c] && !dig2[rc] {
				col[r] = c
				onPath[c], dig1[r+c], dig2[rc] = true, true, true
				dfs(r + 1)
				onPath[c], dig1[r+c], dig2[rc] = false, false, false
			}
		}
	}

	dfs(0)
	return

}

// @lc code=end

/*
// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
