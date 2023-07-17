/*
 * @lc app=leetcode.cn id=2490 lang=golang
 * @lcpr version=21909
 *
 * [2490] 回环句
 */

// @lc code=
package main

func isCircularSentence(sentence string) bool {
	if sentence[0] != sentence[len(sentence)-1] {
		return false
	}
	for i, c := range sentence {
		if c == ' ' && sentence[i-1] != sentence[i+1] {
			return false
		}
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// "leetcode exercises sound delightful"\n
// @lcpr case=end

// @lcpr case=start
// "eetcode"\n
// @lcpr case=end

// @lcpr case=start
// "Leetcode is cool"\n
// @lcpr case=end

*/
