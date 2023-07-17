/*
 * @lc app=leetcode.cn id=60 lang=golang
 * @lcpr version=21909
 *
 * [60] 排列序列
 */

// @lc code=start
// 类似于回溯的爆搜
package main

import (
	"strconv"
	"strings"
)

func getPermutation(n int, k int) string {
	on_path := make([]bool, n+1)
	path := make([]int, n)
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = factorial[i-1] * i
	}

	var dfs func(int, int)
	dfs = func(i, remaining int) {
		if i == n {
			return
		}
		cnt := factorial[n-1-i]
		for j := 1; j < n+1; j++ {
			if !on_path[j] {
				if cnt < remaining {
					remaining -= cnt
					continue
				}
				path[i] = j
				on_path[j] = true
				dfs(i+1, remaining)
				// on_path[j] = false
				return
			}
		}
	}

	dfs(0, k)
	var builder strings.Builder
	for _, num := range path {
		// builder.WriteString(fmt.Sprintf("%d", num))
		builder.WriteString(strconv.Itoa(num))
	}
	return builder.String()
}

// func getPermutation(n int, k int) string {
// 	factorial := make([]int, n)
// 	factorial[0] = 1
// 	for i := 1; i < n; i++ {
// 		factorial[i] = factorial[i-1] * i
// 	}
// 	k--

// 	ans := ""
// 	valid := make([]int, n+1)
// 	for i := 0; i < len(valid); i++ {
// 		valid[i] = 1
// 	}
// 	for i := 1; i <= n; i++ {
// 		order := k/factorial[n-i] + 1
// 		for j := 1; j <= n; j++ {
// 			order -= valid[j]
// 			if order == 0 {
// 				ans += strconv.Itoa(j)
// 				valid[j] = 0
// 				break
// 			}
// 		}
// 		k %= factorial[n-i]
// 	}
// 	return ans
// }

// @lc code=end

/*
// @lcpr case=start
// 3\n3\n
// @lcpr case=end

// @lcpr case=start
// 4\n9\n
// @lcpr case=end

// @lcpr case=start
// 3\n1\n
// @lcpr case=end

*/
