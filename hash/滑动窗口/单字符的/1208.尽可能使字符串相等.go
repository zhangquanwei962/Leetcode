/*
 * @lc app=leetcode.cn id=1208 lang=golang
 * @lcpr version=21908
 *
 * [1208] 尽可能使字符串相等
 */

// @lc code=start
package main

import "fmt"

// 给定两个字符串 s 和 t，以及一个整数 maxCost，计算将 s 中任意子串转换成 t 中对应子串的代价不超过 maxCost 的情况下，s 中最长子串的长度。
func equalSubstring(s string, t string, maxCost int) int {
	left := 0            // 滑动窗口左右端点
	result, cost := 0, 0 // 子串最大长度和当前代价
	fmt.Printf("%T", s[0])
	for right, _ := range s {
		// 在该位置处计算转换代价，并更新当前代价
		cost += abs(int(s[right]) - int(t[right]))

		// 如果当前代价已经超过了最大代价，缩小窗口
		for cost > maxCost {
			cost -= abs(int(s[left]) - int(t[left]))
			left++
		}

		// 记录最长子串的长度
		result = max(result, right-left+1)
	}
	return result
}

// 获取两个数的绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 获取两个数中的最大值
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

/*
// @lcpr case=start
// "abcd"\n"bcdf"\n3\n
// @lcpr case=end

// @lcpr case=start
// "abcd"\n"cdef"\n3\n
// @lcpr case=end

// @lcpr case=start
// "abcd"\n"acde"\n0\n
// @lcpr case=end

*/
