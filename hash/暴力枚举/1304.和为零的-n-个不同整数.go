/*
 * @lc app=leetcode.cn id=1304 lang=golang
 * @lcpr version=21909
 *
 * [1304] 和为零的 N 个不同整数
 */

// @lc code=start
package main

func sumZero(n int) (ans []int) {
	if n&1 == 1 {
		ans = append(ans, 0)
		n--
	}
	for i := 1; i <= n>>1; i++ {
		ans = append(ans, i)
		ans = append(ans, -i)
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// 5\n
// @lcpr case=end

// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
