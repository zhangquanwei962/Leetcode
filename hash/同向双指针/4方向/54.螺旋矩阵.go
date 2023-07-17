/*
 * @lc app=leetcode.cn id=54 lang=golang
 * @lcpr version=21908
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
package main

// type pair struct {
// 	x, y int
// }

// var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func spiralOrder(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	// dx := []int{-1, 0, 1, 0}
	// dy := []int{0, 1, 0, -1}
	dx := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	ans := []int{}

	// 创建二维布尔数组，用于标记已经经过的元素
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	for x, y, i, d := 0, 0, 0, 1; i < n*m; i++ {
		ans = append(ans, matrix[x][y])
		visited[x][y] = true

		// 计算下一个位置的坐标
		// a, b := x+dx[d], y+dy[d]
		a, b := x+dx[d][0], y+dx[d][1]

		// 判断下一个位置是否越界或已经经过
		if a < 0 || a >= n || b < 0 || b >= m || visited[a][b] {
			d = (d + 1) % 4
			a, b = x+dx[d][0], y+dx[d][1]
		}

		// 更新当前位置
		x, y = a, b
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3,4],[5,6,7,8],[9,10,11,12]]\n
// @lcpr case=end

*/
