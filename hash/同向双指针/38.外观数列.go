/*
 * @lc app=leetcode.cn id=38 lang=golang
 * @lcpr version=21909
 *
 * [38] 外观数列
 */

// @lc code=start
package main

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	prev := "1"
	for i := 2; i <= n; i++ {
		cur := &strings.Builder{}
		for j, k := 0, 0; j < len(prev); k = j {
			for j < len(prev) && prev[j] == prev[k] {
				j++
			}
			cur.WriteString(strconv.Itoa(j - k))
			cur.WriteByte(prev[k])
		}
		prev = cur.String()
	}
	return prev
}

// @lc code=end

/*
// @lcpr case=start
// 1\n
// @lcpr case=end

// @lcpr case=start
// 4\n
// @lcpr case=end

*/
