/*
 * @lc app=leetcode.cn id=1017 lang=golang
 * @lcpr version=21907
 *
 * [1017] 负二进制转换
 */

// @lc code=start
package main

import "strconv"

func baseNeg2(n int) string {
	if n == 0 || n == 1 {
		return strconv.Itoa(n)
	}
	res := ""
	for n != 0 {
		remainder := n & 1
		res = strconv.Itoa(remainder) + res
		n -= remainder
		n /= -2
	}

	//	去前导0
	for len(res) > 1 && res[0] == '0' {
		res = res[1:]
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 4\n
// @lcpr case=end

*/
