/*
 * @lc app=leetcode.cn id=2496 lang=golang
 * @lcpr version=21909
 *
 * [2496] 数组中字符串的最大值
 */

// @lc code=start
package main

import "strconv"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maximumValue(strs []string) (res int) {
	for _, s := range strs {
		isdigit := true
		for _, c := range s {
			isdigit = isdigit && (c >= '0') && (c <= '9')
		}

		if isdigit {
			v, _ := strconv.Atoi(s)
			res = max(res, v)
		} else {
			res = max(res, len(s))
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// ["alic3","bob","3","4","00000"]\n
// @lcpr case=end

// @lcpr case=start
// ["1","01","001","0001"]\n
// @lcpr case=end

*/
