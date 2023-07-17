/*
 * @lc app=leetcode.cn id=74 lang=golang
 * @lcpr version=21909
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
// >=能行，<=也能行，找不到就找到-1
package main

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	binary := func(left, right int) int {
		for left+1 < right {
			mid := left + (right-left)>>1
			if matrix[mid/n][mid%n] <= target {
				left = mid
			} else {
				right = mid
			}
		}
		return left
	}
	x := binary(-1, m*n)

	if x == -1 || matrix[x/n][x%n] != target {
		return false
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// [[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n3\n
// @lcpr case=end

// @lcpr case=start
// [[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n13\n
// @lcpr case=end

*/
