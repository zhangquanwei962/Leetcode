/*
 * @lc app=leetcode.cn id=面试题 17.06 lang=golang
 * @lcpr version=21909
 *
 * [面试题 17.06] 2出现的次数
 */

// @lc code=start
package main

import "strconv"

func numberOf2sInRange(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1 // 记忆化
		}
	}

	var f func(int, int, bool) int
	f = func(i, cnt2 int, isLimit bool) (res int) {
		if i == m {
			return cnt2
		}
		if !isLimit {
			dv := &dp[i][cnt2]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}

		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}

		for d := 0; d <= up; d++ {
			c := cnt2
			if d == 2 {
				c++
			}
			res += f(i+1, c, isLimit && d == up)
		}
		return
	}
	return f(0, 0, true)
}

// @lc code=end

/*
// @lcpr case=start
// 25\n
// @lcpr case=end

*/
