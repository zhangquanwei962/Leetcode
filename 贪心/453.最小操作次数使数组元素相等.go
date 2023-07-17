/*
 * @lc app=leetcode.cn id=453 lang=golang
 * @lcpr version=21909
 *
 * [453] 最小操作次数使数组元素相等
 */

// @lc code=start
// 反向思考，n-1个加等于1个减
package main

func minMoves(nums []int) (ans int) {
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	for _, num := range nums {
		ans += num - min
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1]\n
// @lcpr case=end

*/
