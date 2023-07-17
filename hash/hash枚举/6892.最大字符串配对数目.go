/*
 * @lc app=leetcode.cn id=6892 lang=golang
 * @lcpr version=21909
 *
 * [6892] 最大字符串配对数目
 */

// @lc code=start
package main

func maximumNumberOfStringPairs(words []string) (ans int) {
	vis := [26][26]bool{}
	for _, s := range words {
		x, y := s[0]-'a', s[1]-'a'
		if vis[y][x] {
			ans++
		} else {
			vis[x][y] = true
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// ["cd","ac","dc","ca","zz"]\n
// @lcpr case=end

// @lcpr case=start
// ["ab","ba","cc"]\n
// @lcpr case=end

// @lcpr case=start
// ["aa","ab"]\n
// @lcpr case=end

*/
