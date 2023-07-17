/*
 * @lc app=leetcode.cn id=135 lang=golang
 * @lcpr version=21908
 *
 * [135] 分发糖果
 */

// @lc code=start
// 前后缀分解
package main

func candy(ratings []int) (res int) {
	n := len(ratings)
	if n <= 1 {
		return n
	}
	left, right := make([]int, n), make([]int, n)
	for i := range left {
		left[i] = 1
		right[i] = 1
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
	}

	for i := range ratings {
		res += max(left[i], right[i])
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
// [1,0,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,2]\n
// @lcpr case=end

*/
