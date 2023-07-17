/*
 * @lc app=leetcode.cn id=1091 lang=golang
 * @lcpr version=21908
 *
 * [1091] 二进制矩阵中的最短路径
 */

// @lc code=start
package main

type pair struct {
	x, y int
}

var dir8 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	s := pair{0, 0}
	grid[s.x][s.y] = 1

	q := []pair{s}

	for ans := 1; len(q) > 0; ans++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			if p.x == n-1 && p.y == n-1 {
				return ans
			}
			for _, d := range dir8 {
				x, y := p.x+d.x, p.y+d.y
				if 0 <= x && x < n && 0 <= y && y < n && grid[x][y] == 0 {
					// if x == n-1 && y == n-1 {
					// 	return ans
					// }
					grid[x][y] = 1
					q = append(q, pair{x, y})
				}
			}
		}
	}
	return -1

}

// @lc code=end

/*
// @lcpr case=start
// [[0,1],[1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,0,0],[1,1,0],[1,1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,0,0],[1,1,0],[1,1,0]]\n
// @lcpr case=end

*/
