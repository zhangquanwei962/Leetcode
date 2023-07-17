/*
 * @lc app=leetcode.cn id=2606 lang=golang
 * @lcpr version=21909
 *
 * [2606] 找到最大开销的子字符串
 */

// @lc code=start
package main

func maximumCostSubstring(s string, chars string, vals []int) (ans int) {
	mapping := [26]int{}
	for i := range mapping {
		mapping[i] = i + 1
	}
	for i, c := range chars {
		mapping[c-'a'] = vals[i]
	}

	f := 0
	for _, c := range s {
		f = max(f, 0) + mapping[c-'a']
		ans = max(ans, f)
	}
	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// "adaa"\n"d"\n[-1000]\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n"abc"\n[-1,-1,-1]\n
// @lcpr case=end

*/
