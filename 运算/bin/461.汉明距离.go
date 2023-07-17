/*
 * @lc app=leetcode.cn id=461 lang=golang
 * @lcpr version=21909
 *
 * [461] 汉明距离
 */

// @lc code=start
// 异或相同为0，不同为1
package main

func hammingDistance(x int, y int) (ans int) {
	for s := x ^ y; s > 0; s >>= 1 {
		ans += s & 1
	}
	return
}

// func hammingDistance(x, y int) (ans int) {
//     for s := x ^ y; s > 0; s &= s - 1 {
//         ans++
//     }
//     return
// }

// @lc code=end

/*
// @lcpr case=start
// 1\n4\n
// @lcpr case=end

// @lcpr case=start
// 3\n1\n
// @lcpr case=end

*/
