/*
 * @lc app=leetcode.cn id=477 lang=golang
 * @lcpr version=21907
 *
 * [477] 汉明距离总和
 */

// @lc code=start
package main

func totalHammingDistance(nums []int) (ans int) {
	n := len(nums)
	for i := 0; i < 32; i++ {
		c := 0
		for _, x := range nums {
			c += x >> i & 1
		}
		ans += c * (n - c)
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [4,14,2]\n
// @lcpr case=end

// @lcpr case=start
// [4,14,4]\n
// @lcpr case=end

*/
