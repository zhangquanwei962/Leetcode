/*
 * @lc app=leetcode.cn id=1249 lang=golang
 * @lcpr version=21909
 *
 * [1249] 移除无效的括号
 */

// @lc code=start
package main

import "strings"

func minRemoveToMakeValid(s string) string {
	st := []int{}
	vis := make([]bool, len(s))
	for i := range s {
		ch := s[i]
		if ch == byte('(') {
			st = append(st, i)
			vis[i] = true
		} else if ch == byte(')') {
			if len(st) != 0 {
				st = st[:len(st)-1]
				vis[i] = true
			}
		} else {
			vis[i] = true
		}
	}

	for _, v := range st {
		vis[v] = false
	}

	var ans strings.Builder
	for i := 0; i < len(s); i++ {
		if vis[i] {
			ans.WriteByte(s[i])
		}
	}
	return ans.String()
}

// @lc code=end

/*
// @lcpr case=start
// "lee(t(c)o)de)"\n
// @lcpr case=end

// @lcpr case=start
// "a)b(c)d"\n
// @lcpr case=end

// @lcpr case=start
// "))(("\n
// @lcpr case=end

*/
