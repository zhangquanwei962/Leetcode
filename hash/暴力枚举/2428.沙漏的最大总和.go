/*
 * @lc app=leetcode.cn id=2428 lang=golang
 * @lcpr version=21909
 *
 * [2428] 沙漏的最大总和
 */

// @lc code=start
// 3*3卷积
package main

func maxSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	Kernel := [][]int{{1, 1, 1}, {0, 1, 0}, {1, 1, 1}}
	nums := make([][]int, m-2)
	for i := range nums {
		nums[i] = make([]int, n-2)
	}

	for i := 0; i < m-2; i++ {
		for j := 0; j < n-2; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					nums[i][j] += grid[i+k][j+l] * Kernel[k][l]
				}
			}
		}
	}

	maxSum := 0
	for _, row := range nums {
		for _, num := range row {
			if num > maxSum {
				maxSum = num
			}
		}
	}
	return maxSum
}

// @lc code=end

/*
// @lcpr case=start
// [[6,2,1,3],[4,2,1,5],[9,2,8,7],[4,1,2,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

*/
