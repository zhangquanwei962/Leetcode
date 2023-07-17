/*
 * @lc app=leetcode.cn id=338 lang=golang
 * @lcpr version=21909
 *
 * [338] 比特位计数
 */

// @lc code=start
package main

// func lowbit(x int) (ans int) {
// 	for x > 0 {
// 		x = x & (x - 1)
// 		ans++
// 	}
// 	return
// }

func lowbit(x int) (ans int) {
	for x != 0 {
		x -= x & (-x)
		ans++
	}
	return
}

func countBits(n int) (ans []int) {
	for i := 0; i <= n; i++ {
		ans = append(ans, lowbit(i))
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 5\n
// @lcpr case=end

*/
