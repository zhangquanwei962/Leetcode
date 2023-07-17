/*
 * @lc app=leetcode.cn id=2680 lang=golang
 * @lcpr version=21909
 *
 * [2680] 最大或值
 */

// @lc code=start
// 二进制长度要尽量的长
// 不如把k次都乘到一个数上
// 枚举,预处理suf
// 贪心+前后缀分解
package main

func maximumOr(nums []int, k int) int64 {
	n := len(nums)
	suf := make([]int, n+1)

	for i := n - 1; i > 0; i-- {
		suf[i] = suf[i+1] | nums[i]
	}

	ans, pre := 0, 0

	for i, x := range nums {
		ans = max(ans, pre|x<<k|suf[i+1])
		pre |= x
	}
	return int64(ans)

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
// [12,9]\n1\n
// @lcpr case=end

// @lcpr case=start
// [8,1,2]\n2\n
// @lcpr case=end

*/
