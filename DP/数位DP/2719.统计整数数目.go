/*
 * @lc app=leetcode.cn id=2719 lang=golang
 * @lcpr version=21909
 *
 * [2719] 统计整数数目
 */

// @lc code=start
// a-b
// O(10mn) O(mn)
package main

func count(num1, num2 string, minSum, maxSum int) int {
	const mod int = 1e9 + 7
	f := func(s string) int {
		memo := make([][]int, len(s))
		for i := range memo {
			memo[i] = make([]int, min(9*len(s), maxSum)+1)
			for j := range memo[i] {
				memo[i][j] = -1
			}
		}
		var dfs func(p, sum int, limitUp bool) int
		dfs = func(p, sum int, limitUp bool) (res int) {
			if sum > maxSum { // 非法
				return
			}
			if p == len(s) {
				if sum >= minSum { // 合法
					return 1
				}
				return
			}
			if !limitUp {
				ptr := &memo[p][sum]
				if *ptr >= 0 {
					return *ptr
				}
				defer func() { *ptr = res }()
			}
			up := 9
			if limitUp {
				up = int(s[p] - '0')
			}
			for d := 0; d <= up; d++ { // 枚举要填入的数字 d
				res = (res + dfs(p+1, sum+d, limitUp && d == up)) % mod
			}
			return
		}
		return dfs(0, 0, true)
	}
	ans := f(num2) - f(num1) + mod // 避免负数
	sum := 0
	for _, c := range num1 {
		sum += int(c - '0')
	}
	if minSum <= sum && sum <= maxSum { // x=num1 是合法的，补回来
		ans++
	}
	return ans % mod
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// "1"\n"12"\n1\n8\n
// @lcpr case=end

// @lcpr case=start
// "1"\n"5"\n1\n5\n
// @lcpr case=end

*/
