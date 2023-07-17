/*
 * @lc app=leetcode.cn id=1738 lang=golang
 * @lcpr version=21909
 *
 * [1738] 找出第 K 大的异或坐标值
 */

// @lc code=start
package main

import "math/rand"

func quickSelect(a []int, k int) int {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	for l, r := 0, len(a)-1; l < r; {
		v := a[l]
		i, j := l, r+1
		for {
			for i++; i < r && a[i] < v; i++ {
			}
			for j--; j > l && a[j] > v; j-- {
			}
			if i >= j {
				break
			}
			a[i], a[j] = a[j], a[i]
		}
		a[l], a[j] = a[j], v
		if j == k {
			break
		} else if j < k {
			l = j + 1
		} else {
			r = j - 1
		}
	}
	return a[k]
}

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	results := make([]int, 0, m*n)
	pre := make([][]int, m+1)
	pre[0] = make([]int, n+1)
	for i, row := range matrix {
		pre[i+1] = make([]int, n+1)
		for j, val := range row {
			pre[i+1][j+1] = pre[i+1][j] ^ pre[i][j+1] ^ pre[i][j] ^ val
			results = append(results, pre[i+1][j+1])
		}
	}
	return quickSelect(results, m*n-k)
}

// @lc code=end

/*
// @lcpr case=start
// [[5,2],[1,6]]\n1\n
// @lcpr case=end

// @lcpr case=start
// [[5,2],[1,6]]\n2\n
// @lcpr case=end

// @lcpr case=start
// [[5,2],[1,6]]\n3\n
// @lcpr case=end

// @lcpr case=start
// [[5,2],[1,6]]\n4\n
// @lcpr case=end

*/
