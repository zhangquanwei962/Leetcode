/*
 * @lc app=leetcode.cn id=2711 lang=golang
 * @lcpr version=21909
 *
 * [2711] 对角线上不同值的数量差
 */

// @lc code=start
// O(mn) O(mn)
package main

func differenceOfDistinctValues(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for s := 1; s < m+n; s++ { // 编号
		minJ := max(0, n-s)       // 列编号最小值
		maxJ := min(n-1, n-s+m-1) // 列编号最大值
		// topLeft
		set := map[int]struct{}{}
		for j := minJ; j < maxJ; j++ {
			i := s + j - n // 行编号
			set[grid[i][j]] = struct{}{}
			ans[i+1][j+1] = len(set)
		}
		// bottomRight
		set = map[int]struct{}{}
		for j := maxJ; j > minJ; j-- {
			i := s + j - n
			set[grid[i][j]] = struct{}{}
			ans[i-1][j-1] = abs(ans[i-1][j-1] - len(set))
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// func differenceOfDistinctValues(grid [][]int) [][]int {
// 	m, n := len(grid), len(grid[0])
// 	ans := make([][]int, m)
// 	for i := range ans {
// 		ans[i] = make([]int, n)
// 		for j := range ans[i] {
// 			// topLeft
// 			set := map[int]struct{}{}
// 			for x, y := i-1, j-1; x >= 0 && y >= 0; {
// 				set[grid[x][y]] = struct{}{}
// 				x--
// 				y--
// 			}
// 			sz := len(set)

// 			// bottomRight
// 			set = map[int]struct{}{}
// 			for x, y := i+1, j+1; x < m && y < n; {
// 				set[grid[x][y]] = struct{}{}
// 				x++
// 				y++
// 			}
// 			ans[i][j] = abs(sz - len(set))
// 		}
// 	}
// 	return ans
// }

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }

// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[3,1,5],[3,2,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1]]\n
// @lcpr case=end

*/
