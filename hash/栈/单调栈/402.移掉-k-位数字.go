/*
 * @lc app=leetcode.cn id=402 lang=golang
 * @lcpr version=21909
 *
 * [402] 移掉 K 位数字
 */

// @lc code=start
// 单调栈+贪心
package main

import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	// n := len(num)
	stk := []byte{}
	for i := range num {
		digit := num[i] // 单调不降
		for k > 0 && len(stk) > 0 && digit < stk[len(stk)-1] {
			stk = stk[:len(stk)-1]
			k--
		}
		stk = append(stk, digit)
	}
	fmt.Println(stk)
	// 如果单调递增，移除后面比较大的
	stk = stk[:len(stk)-k]
	ans := strings.TrimLeft(string(stk), "0")
	if ans == "" {
		return "0"
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// "1432219"\n3\n
// @lcpr case=end

// @lcpr case=start
// "10200"\n1\n
// @lcpr case=end

// @lcpr case=start
// "10"\n2\n
// @lcpr case=end

*/
