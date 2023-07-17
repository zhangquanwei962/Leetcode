/*
 * @lc app=leetcode.cn id=494 lang=golang
 * @lcpr version=21907
 *
 * [494] 目标和
 */

// @lc code=start

// 如果不初始化的话，需要+1的
// 至多是什么？初始化全1，取消掉c==0判断
// 至少是什么？循环到0，max(j-x,0)
package main

func findTargetSumWays(nums []int, target int) int {
	for _, x := range nums {
		target += x
	}

	if target < 0 || target&1 == 1 {
		return 0
	}
	target /= 2
	// n := len(nums)

	f := make([]int, target+1)
	f[0] = 1

	for _, x := range nums {
		for j := target; j >= x; j-- {
			f[j] += f[j-x]
		}
	}
	return f[target]
}

// func findTargetSumWays(nums []int, target int) int {
// 	for _, x := range nums {
// 		target += x
// 	}

// 	if target < 0 || target&1 == 1 {
// 		return 0
// 	}
// 	target /= 2
// 	n := len(nums)

// 	f := make([][]int, n+1)
// 	for i := range f {
// 		f[i] = make([]int, target+1)
// 	}

// 	f[0][0] = 1
// 	for i, x := range nums {
// 		for c := 0; c <= target; c++ {
// 			if c < x {
// 				f[i+1][c] = f[i][c]
// 			} else {
// 				f[i+1][c] = f[i][c] + f[i][c-x]
// 			}
// 		}
// 	}
// 	return f[n][target]

// var dfs func(int, int) int
// dfs = func(i, c int) int {
// 	if i < 0 {
// 		if c == 0 {
// 			return 1
// 		}
// 		return 0
// 	}
// 	if c < nums[i] {
// 		return dfs(i-1, c)
// 	}
// 	return dfs(i-1, c) + dfs(i-1, c-nums[i])
// }
// return dfs(n-1, target)
// }

// @lc code=end

/*
// @lcpr case=start
// [1,1,1,1,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

*/
