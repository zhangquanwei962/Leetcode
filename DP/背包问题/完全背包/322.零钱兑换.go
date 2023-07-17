/*
 * @lc app=leetcode.cn id=322 lang=golang
 * @lcpr version=21908
 *
 * [322] 零钱兑换
 */

// @lc code=start
package main

import "math"

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func coinChange(coins []int, amount int) int {
	n := len(coins)

	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, amount+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int) int

	dfs = func(i, c int) (res int) {
		if i < 0 {
			if c == 0 {
				return 0
			}
			return math.MaxInt / 2
		}

		C := &memo[i][c]
		if *C != -1 {
			return *C
		}
		defer func() {
			*C = res
		}()

		if c < coins[i] {
			return dfs(i-1, c)
		}
		return min(dfs(i-1, c), dfs(i, c-coins[i])+1)
	}
	ans := dfs(n-1, amount)
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}

// f[i]代表组成金额i所需要的最少金币个数

// func coinChange(coins []int, amount int) int {
//     f := make([]int, amount+1)
//     for i := range f {
//         f[i] = math.MaxInt / 2 // 除 2 是防止下面 + 1 溢出
//     }
//     f[0] = 0
//     for _, x := range coins {
//         for c := x; c <= amount; c++ {
//             f[c] = min(f[c], f[c-x]+1)
//         }
//     }
//     ans := f[amount]
//     if ans < math.MaxInt/2 {
//         return ans
//     }
//     return -1
// }

// func min(a, b int) int { if a > b { return b }; return a }

// @lc code=end

/*
// @lcpr case=start
// [1, 2, 5]\n11\n
// @lcpr case=end

// @lcpr case=start
// [2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n0\n
// @lcpr case=end

*/
