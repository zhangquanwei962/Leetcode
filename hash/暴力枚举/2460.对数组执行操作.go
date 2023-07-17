/*
 * @lc app=leetcode.cn id=2460 lang=golang
 * @lcpr version=21908
 *
 * [2460] 对数组执行操作
 */

// @lc code=start
package main

func applyOperations(a []int) []int {
	n := len(a)
	b := a[:0]
	for i := 0; i < n-1; i++ {
		if a[i] > 0 {
			if a[i] == a[i+1] {
				a[i] *= 2
				a[i+1] = 0
			}
			b = append(b, a[i])
		}
	}

	if a[n-1] > 0 {
		b = append(b, a[n-1])
	}

	for i := len(b); i < n; i++ {
		a[i] = 0
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,2,1,1,0]\n
// @lcpr case=end

// @lcpr case=start
// [0,1]\n
// @lcpr case=end

*/
