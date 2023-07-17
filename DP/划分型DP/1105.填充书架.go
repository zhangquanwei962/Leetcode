/*
 * @lc app=leetcode.cn id=1105 lang=golang
 * @lcpr version=21907
 *
 * [1105] 填充书架
 */

// @lc code=start
package main

import "math"

func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	f := make([]int, n+1) // f[0]=0，翻译自 dfs(-1)=0
	for i := range books {
		f[i+1] = math.MaxInt
		maxH, leftW := 0, shelfWidth
		for j := i; j >= 0; j-- {
			leftW -= books[j][0]
			if leftW < 0 { // 空间不足，无法放书
				break
			}
			maxH = max(maxH, books[j][1]) // 从 j 到 i 的最大高度
			f[i+1] = min(f[i+1], f[j]+maxH)
		}
	}
	return f[n] // 翻译自 dfs(n-1)
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

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// func min(x, y int) int {
// 	if x > y {
// 		return y
// 	}
// 	return x
// }
// func minHeightShelves(books [][]int, shelfWidth int) int {
// 	n := len(books)
// 	memo := make([]int, n)
// 	for i := range memo {
// 		memo[i] = -1
// 	}

// 	var dfs func(int) int

// 	dfs = func(i int) int {
// 		if i < 0 {
// 			return 0
// 		}
// 		p := &memo[i]
// 		if *p != -1 {
// 			return *p
// 		}

// 		res := math.MaxInt

// 		defer func() {
// 			*p = res
// 		}()

// 		maxH, leftW := 0, shelfWidth
// 		for j := i; j >= 0; j-- {
// 			leftW -= books[j][0]
// 			if leftW < 0 {
// 				break
// 			}

// 			maxH = max(maxH, books[j][1]) // 从 j 到 i 的最大高度
// 			res = min(res, dfs(j-1)+maxH)
// 		}
// 		return res

// 	}
// 	return dfs(n - 1)
// }

// @lc code=end

/*
// @lcpr case=start
// [[1,1],[2,3],[2,3],[1,1],[1,1],[1,1],[1,2]]\n4\n
// @lcpr case=end

// @lcpr case=start
// [[1,3],[2,4],[3,2]]\n6\n
// @lcpr case=end

*/
