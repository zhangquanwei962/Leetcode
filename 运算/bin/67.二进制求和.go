/*
 * @lc app=leetcode.cn id=67 lang=golang
 * @lcpr version=21907
 *
 * [67] 二进制求和
 */

// @lc code=start
package main

import "strconv"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func addBinary(a string, b string) string {
	i1, i2 := len(a)-1, len(b)-1
	carry, res := 0, ""

	for i1 >= 0 || i2 >= 0 || carry != 0 {
		if i1 >= 0 {
			num, _ := strconv.Atoi(string(a[i1]))
			carry += num
			i1--
		}
		if i2 >= 0 {
			carry += int(b[i2] - '0')
			i2--
		}
		res = strconv.Itoa(carry&1) + res
		carry >>= 1
	}

	for len(res) > 1 && res[0] == '0' {
		res = res[1:]
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// "11"\n"1"\n
// @lcpr case=end

// @lcpr case=start
// "1010"\n"1011"\n
// @lcpr case=end

*/
