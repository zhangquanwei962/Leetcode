/*
 * @lc app=leetcode.cn id=2370 lang=golang
 * @lcpr version=21907
 *
 * [2370] 最长理想子序列
 */

// @lc code=start
package main

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func longestIdealString(s string, k int) (res int) {
	f := [26]int{}
	for _, c := range s {
		c := int(c - 'a')
		for _, v := range f[max(c-k, 0):min(c+k+1, 26)] {
			f[c] = max(f[c], v)
		}
		f[c]++
	}

	for _, v := range f {
		res = max(res, v)
	}
	return

}

// @lc code=end

/*
// @lcpr case=start
// "acfgbd"\n2\n
// @lcpr case=end

// @lcpr case=start
// "abcd"\n3\n
// @lcpr case=end

*/
