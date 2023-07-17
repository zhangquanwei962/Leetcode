/*
 * @lc app=leetcode.cn id=2437 lang=golang
 * @lcpr version=21908
 *
 * [2437] 有效时间的数目
 */

// @lc code=start

// 这段代码通过拆分时间字符串，并遍历可选数字来计算出可以组成合法时间的方案数
// ，从而实现了计算“可行方案数”的功能
package main

func countTime(time string) int {
	f := func(s string, m int) (cnt int) {
		for i := 0; i < m; i++ {
			a := s[0] == '?' || int(s[0]-'0') == i/10
			b := s[1] == '?' || int(s[1]-'0') == i%10
			if a && b {
				cnt++
			}
		}
		return
	}
	return f(time[:2], 24) * f(time[3:], 60)
}

// @lc code=end

/*
// @lcpr case=start
// "?5:00"\n
// @lcpr case=end

// @lcpr case=start
// "0?:0?"\n
// @lcpr case=end

// @lcpr case=start
// "??:??"\n
// @lcpr case=end

*/
