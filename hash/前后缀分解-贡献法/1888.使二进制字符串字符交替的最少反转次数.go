/*
 * @lc app=leetcode.cn id=1888 lang=golang
 * @lcpr version=21909
 *
 * [1888] 使二进制字符串字符交替的最少反转次数
 */

// @lc code=start
// 数据范围10^5,O(n^2)超时，考虑到1操作就是首尾相连字符串节
// 各个字符串相对位置不变，那么其实就是1010... 0101...循环节
// 长度为奇数的话，就会多出来11或者00
// 可以枚举其出现位置，对该位置左边按照 1 开头的交替字符串来匹配，右边按照
// 0 结尾的交替字符串来匹配；或者左边按照 0 开头的交替字符串来匹配，右边按照
// 1 结尾的交替字符串来匹配。 O(n) O(n)
package main

func minFlips(s string) int {
	n := len(s)
	ans := n
	// 枚举开头是 0 还是 1
	for head := byte('0'); head <= '1'; head++ {
		// 左边每个位置的不同字母个数
		leftDiff := make([]int, n)
		diff := 0
		for i := range s {
			if s[i] != head^byte(i&1) {
				diff++
			}
			leftDiff[i] = diff
		}
		// 右边每个位置的不同字母个数
		tail := head ^ 1
		diff = 0
		for i := n - 1; i >= 0; i-- {
			// 左边+右边即为整个字符串的不同字母个数，取最小值
			ans = min(ans, leftDiff[i]+diff)
			if s[i] != tail^byte((n-1-i)&1) {
				diff++
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// "111000"\n
// @lcpr case=end

// @lcpr case=start
// "010"\n
// @lcpr case=end

// @lcpr case=start
// "1110"\n
// @lcpr case=end

*/
