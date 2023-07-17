/*
 * @lc app=leetcode.cn id=71 lang=golang
 * @lcpr version=21909
 *
 * [71] 简化路径
 */

// @lc code=start
package main

import (
	"strings"
)

func simplifyPath(path string) string {
	stk := []string{}
	for _, name := range strings.Split(path, "/") {
		if name == ".." {
			if len(stk) > 0 {
				stk = stk[:len(stk)-1]
			}
		} else if name != "" && name != "." {
			stk = append(stk, name)
		}
	}
	return "/" + strings.Join(stk, "/")
}

// @lc code=end

/*
// @lcpr case=start
// "/home/"\n
// @lcpr case=end

// @lcpr case=start
// "/../"\n
// @lcpr case=end

// @lcpr case=start
// "/home//foo/"\n
// @lcpr case=end

// @lcpr case=start
// "/a/./b/../../c/"\n
// @lcpr case=end

*/
