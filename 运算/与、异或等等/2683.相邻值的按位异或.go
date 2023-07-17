/*
 * @lc app=leetcode.cn id=2683 lang=golang
 * @lcpr version=21909
 *
 * [2683] 相邻值的按位异或
 */

// @lc code=start
package main

func doesValidArrayExist(derived []int) bool {
	xor := 0
	for _, x := range derived {
		xor ^= x
	}
	return xor == 0
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,0]\n
// @lcpr case=end

// @lcpr case=start
// [1,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,0]\n
// @lcpr case=end

*/
