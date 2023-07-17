/*
 * @lc app=leetcode.cn id=2575 lang=golang
 * @lcpr version=21909
 *
 * [2575] 找出字符串的可整除数组
 */

// @lc code=start
// (a+b)mod P = (a mod P+b mod P)mod P
// 小数论，若a % b = c， 则ka % b = kc % b
package main

func divisibilityArray(word string, m int) []int {
	ans := make([]int, len(word))
	x := 0
	for i, c := range word {
		x = (x*10 + int(c-'0')) % m
		if x == 0 {
			ans[i] = 1
		}
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// "998244353"\n3\n
// @lcpr case=end

// @lcpr case=start
// "1010"\n10\n
// @lcpr case=end

*/
