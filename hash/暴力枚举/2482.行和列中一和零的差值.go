/*
 * @lc app=leetcode.cn id=2482 lang=golang
 * @lcpr version=21909
 *
 * [2482] 行和列中一和零的差值
 */

// @lc code=start
// 行和列都可以看成1的个数减去0的个数
// 于是x*2-1
package main

func onesMinusZeros(grid [][]int) [][]int {
	r := make([]int, len(grid))
	c := make([]int, len(grid[0]))
	for i, row := range grid {
		for j, x := range row {
			r[i] += x*2 - 1
			c[j] += x*2 - 1 // 1 -> 1, 0 -> -1
		}
	}
	for i, x := range r {
		for j, y := range c {
			grid[i][j] = x + y
		}
	}
	return grid
}

// @lc code=end

/*
// @lcpr case=start
// [[0,1,1],[1,0,1],[0,0,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,1,1],[1,1,1]]\n
// @lcpr case=end

*/
