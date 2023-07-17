/*
 * @lc app=leetcode.cn id=415 lang=golang
 * @lcpr version=21909
 *
 * [415] 字符串相加
 */

// @lc code=start
package main

import "strconv"

func addStrings(num1 string, num2 string) (res string) {
	m, n := len(num1), len(num2)
	i, j, carry := m-1, n-1, 0
	for i >= 0 || j >= 0 || carry != 0 {
		if i >= 0 {
			carry += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			carry += int(num2[j] - '0')
			j--
		}
		res = strconv.Itoa(carry%10) + res
		carry /= 10
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// "11"\n"123"\n
// @lcpr case=end

// @lcpr case=start
// "456"\n"77"\n
// @lcpr case=end

// @lcpr case=start
// "0"\n"0"\n
// @lcpr case=end

*/
