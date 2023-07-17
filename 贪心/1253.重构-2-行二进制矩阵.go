/*
 * @lc app=leetcode.cn id=1253 lang=golang
 * @lcpr version=21909
 *
 * [1253] 重构 2 行二进制矩阵
 */

// @lc code=start
package main

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	n := len(colsum)
	ans := make([][]int, 2)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for j, v := range colsum {
		if v == 2 {
			ans[0][j], ans[1][j] = 1, 1
			upper--
			lower--
		}

		if v == 1 {
			if upper > lower {
				upper--
				ans[0][j] = 1
			} else {
				lower--
				ans[1][j] = 1
			}
		}

		if upper < 0 || lower < 0 {
			break
		}
	}
	if upper != 0 || lower != 0 {
		return [][]int{}
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// 2\n1\n[1,1,1]\n
// @lcpr case=end

// @lcpr case=start
// 2\n3\n[2,2,1,1]\n
// @lcpr case=end

// @lcpr case=start
// 5\n5\n[2,1,2,0,1,0,1,2,0,1]\n
// @lcpr case=end

*/
