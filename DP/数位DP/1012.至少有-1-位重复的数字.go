/*
 * @lc app=leetcode.cn id=1012 lang=golang
 * @lcpr version=21909
 *
 * [1012] 至少有 1 位重复的数字
 */

// @lc code=start
package main

import "strconv"

func numDupDigitsAtMostN(n int) (ans int) {
	s := strconv.Itoa(n)
	m := len(s)
	memo := make([][1 << 10]int, m)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var f func(int, int, bool, bool) int
	f = func(i, mask int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum {
				return 1 // 得到了一个合法数字
			}
			return
		}
		if !isLimit && isNum {
			p := &memo[i][mask]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
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
			up = int(s[i] - '0') // 如果前面填的数字都和 n 的一样，那么这一位至多填数字 s[i]（否则就超过 n 啦）
		}
		for ; d <= up; d++ { // 枚举要填入的数字 d
			if mask>>d&1 == 0 { // d 不在 mask 中
				res += f(i+1, mask|1<<d, isLimit && d == up, true)
			}
		}
		return
	}
	return n - f(0, 0, true, false)
}

// @lc code=end

/*
// @lcpr case=start
// 20\n
// @lcpr case=end

// @lcpr case=start
// 100\n
// @lcpr case=end

// @lcpr case=start
// 1000\n
// @lcpr case=end

*/
