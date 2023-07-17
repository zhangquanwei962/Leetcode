/*
 * @lc app=leetcode.cn id=440 lang=golang
 * @lcpr version=21909
 *
 * [440] 字典序的第K小数字
 */

// @lc code=start
// 字节hard TOP10
// log^2(n) O(1)
package main

func getSteps(cur, n int) (steps int) {
	first, last := cur, cur
	for first <= n {
		steps += min(last, n) - first + 1
		first *= 10
		last = last*10 + 9
	}
	return
}

func findKthNumber(n, k int) int {
	cur := 1
	k--
	for k > 0 {
		steps := getSteps(cur, n)
		if steps <= k {
			k -= steps
			cur++
		} else {
			cur *= 10
			k--
		}
	}
	return cur
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// 13\n2\n
// @lcpr case=end

// @lcpr case=start
// 1\n1\n
// @lcpr case=end

*/
