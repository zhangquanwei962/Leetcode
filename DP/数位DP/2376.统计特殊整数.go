/*
 * @lc app=leetcode.cn id=2376 lang=golang
 * @lcpr version=21909
 *
 * [2376] 统计特殊整数
 */

// @lc code=start
package main

import "strconv"

func countSpecialNumbers(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	dp := make([][1 << 10]int, m)
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1 // 记忆化
		}
	}

	var f func(int, int, bool, bool) int
	f = func(i, mask int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum {
				return 1
			}
			return
		}
		if !isLimit && isNum {
			dv := &dp[i][mask]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}
		if !isNum { // 可以跳过当前数位
			res += f(i+1, mask, false, false)
		}
		d := 0
		if !isNum {
			d = 1 // 如果前面没有填数字，必须从 1 开始（因为不能有前导零）
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}

		for ; d <= up; d++ {
			if mask>>d&1 == 0 {
				res += f(i+1, mask|1<<d, isLimit && d == up, true)
			}
		}
		return
	}
	return f(0, 0, true, false)
}

// @lc code=end

/*
// @lcpr case=start
// 20\n
// @lcpr case=end

// @lcpr case=start
// 5\n
// @lcpr case=end

// @lcpr case=start
// 135\n
// @lcpr case=end

*/
