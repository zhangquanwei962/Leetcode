/*
 * @lc app=leetcode.cn id=59 lang=golang
 * @lcpr version=21909
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
package main

type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // 左上右下

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for x, y, i, d := 0, 0, 1, 1; i <= n*n; i++ {
		matrix[x][y] = i
		a, b := x+dirs[d].x, y+dirs[d].y
		if a < 0 || a >= n || b < 0 || b >= n || matrix[a][b] != 0 {
			d = (d + 1) % 4
			a, b = x+dirs[d].x, y+dirs[d].y
		}
		x, y = a, b
	}
	return matrix
}

// @lc code=end

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
