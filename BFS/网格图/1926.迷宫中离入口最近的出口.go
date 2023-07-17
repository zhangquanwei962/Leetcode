/*
 * @lc app=leetcode.cn id=1926 lang=golang
 * @lcpr version=21908
 *
 * [1926] 迷宫中离入口最近的出口
 */

// @lc code=start
// 入口不算出口
package main

func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	q := []pair{{entrance[0], entrance[1]}}
	maze[entrance[0]][entrance[1]] = '+'

	for ans := 1; len(q) > 0; ans++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			X, Y := p.x, p.y
			// if X == 0 || X == m-1 || Y == 0 || Y == n-1 {
			// 	return ans
			// }
			for _, d := range dirs4 {
				x, y := X+d[0], Y+d[1]
				if 0 <= x && x < m && 0 <= y && y < n && maze[x][y] == '.' {
					if x == 0 || x == m-1 || y == 0 || y == n-1 {
						return ans
					}
					maze[x][y] = '+'
					q = append(q, pair{x, y})
				}
			}
		}
	}
	return -1
}

type pair struct{ x, y int }

// var dx = []int{0, 1, 0, -1}
// var dy = []int{-1, 0, 1, 0}
var dirs4 = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// @lc code=end

/*
// @lcpr case=start
// [["+","+",".","+"],[".",".",".","+"],["+","+","+","."]]\n[1,2]\n
// @lcpr case=end

// @lcpr case=start
// [["+","+","+"],[".",".","."],["+","+","+"]]\n[1,0]\n
// @lcpr case=end

// @lcpr case=start
// [[".","+"]]\n[0,0]\n
// @lcpr case=end

*/
