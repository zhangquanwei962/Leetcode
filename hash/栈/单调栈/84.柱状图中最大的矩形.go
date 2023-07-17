/*
 * @lc app=leetcode.cn id=84 lang=golang
 * @lcpr version=21909
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
// 左侧严格小于 右侧小于等于
// O(n) O(n)
package main

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	mono_stack := []int{-1}
	for i := 0; i < n; i++ {
		for len(mono_stack) > 1 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			right[mono_stack[len(mono_stack)-1]] = i
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		left[i] = mono_stack[len(mono_stack)-1]
		mono_stack = append(mono_stack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

/*
// @lcpr case=start
// [2,1,5,6,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [2,4]\n
// @lcpr case=end

*/
