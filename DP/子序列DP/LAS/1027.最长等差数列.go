/*
 * @lc app=leetcode.cn id=1027 lang=golang
 * @lcpr version=21909
 *
 * [1027] 最长等差数列
 */

// @lc code=start
package main

// 常数优化
func longestArithSeqLength(a []int) (ans int) {
	n := len(a)
	f := make([][1001]int, n)
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			d := a[i] - a[j] + 500 // +500 防止出现负数
			if f[i][d] == 0 {
				f[i][d] = f[j][d] + 1 // 默认的 1 在下面返回时加上
				ans = max(ans, f[i][d])
			}
		}
	}
	return ans + 1
}

// dp[i][d]表示以nums[i]结尾且公差为d的数列长度。
// func longestArithSeqLength(a []int) (ans int) {
// 	n := len(a)
// 	f := make([]map[int]int, n)
// 	f[0] = map[int]int{}

// 	for i := 1; i < n; i++ {
// 		f[i] = map[int]int{}
// 		for j := i - 1; j >= 0; j-- {
// 			d := a[i] - a[j]
// 			if f[i][d] == 0 {
// 				f[i][d] = f[j][d] + 1
// 				ans = max(ans, f[i][d])
// 			}
// 		}
// 	}

// 	return ans + 1
// }

// ** 记忆化搜索 **
// func longestArithSeqLength(a []int) (ans int) {
// 	n := len(a)
// 	maxLen := make([]map[int]int, n)
// 	var dfs func(int) map[int]int
// 	dfs = func(i int) map[int]int {
// 		if maxLen[i] != nil { // 之前算过了
// 			return maxLen[i]
// 		}
// 		maxLen[i] = map[int]int{}
// 		for j := i - 1; j >= 0; j-- {
// 			d := a[i] - a[j] // 公差
// 			if maxLen[i][d] == 0 {
// 				maxLen[i][d] = dfs(j)[d] + 1
// 				ans = max(ans, maxLen[i][d])
// 			}
// 		}
// 		return maxLen[i]
// 	}
// 	for i := 1; i < n; i++ {
// 		dfs(i)
// 	}
// 	return ans + 1
// }

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [3,6,9,12]\n
// @lcpr case=end

// @lcpr case=start
// [9,4,7,2,10]\n
// @lcpr case=end

// @lcpr case=start
// [20,1,15,3,10,5,8]\n
// @lcpr case=end

*/
