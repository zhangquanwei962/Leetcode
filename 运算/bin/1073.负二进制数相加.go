/*
 * @lc app=leetcode.cn id=1073 lang=golang
 * @lcpr version=21907
 *
 * [1073] 负二进制数相加
 */

// @lc code=start
package main

func addNegabinary(arr1 []int, arr2 []int) []int {
	i1, i2 := len(arr1)-1, len(arr2)-1
	carry, res := 0, []int{}
	for i1 >= 0 || i2 >= 0 || carry != 0 {
		if i1 >= 0 {
			carry += arr1[i1]
			i1--
		}
		if i2 >= 0 {
			carry += arr2[i2]
			i2--
		}
		// 由于我们使用的是基数为 -2 的二进制数，
		// 所以最低位只可能是 1 或 0
		res = append(res, carry&1)
		carry = -(carry >> 1)
	}
	for len(res) > 1 && res[len(res)-1] == 0 {
		res = res[:len(res)-1]
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,1,1,1]\n[1,0,1]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[1]\n
// @lcpr case=end

*/
