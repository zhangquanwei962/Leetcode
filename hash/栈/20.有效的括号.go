/*
 * @lc app=leetcode.cn id=20 lang=golang
 * @lcpr version=21909
 *
 * [20] 有效的括号
 */

// @lc code=start
package main

func isValid(s string) bool {
	if len(s)&1 == 1 {
		return false
	}

	pairs := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	st := []string{}

	for i := range s {
		ch := string(s[i])
		if _, ok := pairs[ch]; ok {
			if len(st) == 0 || st[len(st)-1] != pairs[ch] {
				return false
			}
			st = st[:len(st)-1]
		} else {
			st = append(st, ch)
		}
	}
	if len(st) == 0 {
		return true
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// "()"\n
// @lcpr case=end

// @lcpr case=start
// "()[]{}"\n
// @lcpr case=end

// @lcpr case=start
// "(]"\n
// @lcpr case=end

*/
