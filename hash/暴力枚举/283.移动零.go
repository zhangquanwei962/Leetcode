/*
 * @lc app=leetcode.cn id=283 lang=golang
 * @lcpr version=21908
 *
 * [283] 移动零
 */

// @lc code=start
// 1在0前面，引用
package main

func moveZeroes(a []int) {
	n := len(a)
	b := a[:0]
	for i := 0; i < n; i++ {
		if a[i] != 0 {
			b = append(b, a[i])
		}
	}

	for i := len(b); i < n; i++ {
		a[i] = 0
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [0,1,0,3,12]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/
