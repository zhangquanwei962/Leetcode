/*
 * @lc app=leetcode.cn id=2527 lang=golang
 * @lcpr version=21909
 *
 * [2527] 查询数组 Xor 美丽值
 */

// @lc code=start
// 位运算经典技巧：由于每个比特位互不相干，所以拆分成每个比特位分别计算。
// 由于对异或影响的只有1，只需要计算1的个数，
// 异或又叫做不进位加法
package main

func xorBeauty(nums []int) (ans int) {
	for _, x := range nums {
		ans ^= x
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [1,4]\n
// @lcpr case=end

// @lcpr case=start
// [15,45,20,2,34,35,5,44,32,30]\n
// @lcpr case=end

*/
