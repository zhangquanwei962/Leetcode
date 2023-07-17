/*
 * @lc app=leetcode.cn id=204 lang=golang
 * @lcpr version=21909
 *
 * [204] 计数质数
 */

// @lc code=start
package main

func countPrimes(n int) (ans int) {
	np := make([]bool, n)
	// var p []int
	for i := 2; i < n; i++ {
		if !np[i] {
			// p = append(p, i) // 预处理质数列表
			ans++
			for j := i * i; j < n; j += i {
				np[j] = true
			}
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// 10\n
// @lcpr case=end

// @lcpr case=start
// 0\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
