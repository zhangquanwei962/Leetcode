/*
 * @lc app=leetcode.cn id=137 lang=golang
 * @lcpr version=21907
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start

/*
加起来二进制之和
*/
package main

func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}

// @lc code=end

/*
// @lcpr case=start
// [2,2,3,2]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,0,1,0,1,99]\n
// @lcpr case=end

*/
