/*
 * @lc app=leetcode.cn id=1375 lang=golang
 * @lcpr version=21909
 *
 * [1375] 二进制字符串前缀一致的次数
 */

// @lc code=start
package main

func numTimesAllBlue(flips []int) (ans int) {
	mx := 0
	for i, x := range flips {
		mx = max(x, mx)
		if mx == i+1 {
			ans++
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [3,2,4,1,5]\n
// @lcpr case=end

// @lcpr case=start
// [4,1,2,3]\n
// @lcpr case=end

*/
