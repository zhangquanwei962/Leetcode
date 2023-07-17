/*
 * @lc app=leetcode.cn id=1137 lang=golang
 * @lcpr version=21909
 *
 * [1137] 第 N 个泰波那契数
 */

// @lc code=start
package main

var p [38]int

func init() {
	p[0], p[1], p[2] = 0, 1, 1
	for i := 3; i < 38; i++ {
		p[i] = p[i-1] + p[i-2] + p[i-3]
	}
}
func tribonacci(n int) int {
	return p[n]
}

// @lc code=end

/*
// @lcpr case=start
// 4\n
// @lcpr case=end

// @lcpr case=start
// 25\n
// @lcpr case=end

*/
