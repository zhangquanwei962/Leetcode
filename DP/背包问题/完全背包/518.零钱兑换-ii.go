/*
 * @lc app=leetcode.cn id=518 lang=golang
 * @lcpr version=21908
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
// f[i]代表组成金额i有多少种方案
package main

func change(amount int, coins []int) int {
	f := make([]int, amount+1)
	f[0] = 1
	for _, x := range coins {
		for j := x; j <= amount; j++ {
			f[j] += f[j-x]
		}
	}
	return f[amount]

}

// @lc code=end

/*
// @lcpr case=start
// 5\n[1, 2, 5]\n
// @lcpr case=end

// @lcpr case=start
// 3\n[2]\n
// @lcpr case=end

// @lcpr case=start
// 10\n[10]\n
// @lcpr case=end

*/
