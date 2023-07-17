/*
 * @lc app=leetcode.cn id=1419 lang=golang
 * @lcpr version=21907
 *
 * [1419] 数青蛙
 */

// @lc code=start
package main

import "fmt"

var previous = [...]int{'c': 'k', 'r': 'c', 'o': 'r', 'a': 'o', 'k': 'a'}

func minNumberOfFrogs(croakOfFrogs string) int {
	cnt := [len(previous)]int{}
	fmt.Println(previous)
	for _, ch := range croakOfFrogs {
		pre := previous[ch]
		if cnt[pre] > 0 {
			cnt[pre]--
		} else if ch != 'c' {
			return -1
		}
		cnt[ch]++
	}

	if cnt['c'] > 0 || cnt['r'] > 0 || cnt['o'] > 0 || cnt['a'] > 0 {
		return -1
	}

	return cnt['k']
}

// @lc code=end

/*
// @lcpr case=start
// "croakcroak"\n
// @lcpr case=end

// @lcpr case=start
// "crcoakroak"\n
// @lcpr case=end

// @lcpr case=start
// "croakcrook"\n
// @lcpr case=end

*/
