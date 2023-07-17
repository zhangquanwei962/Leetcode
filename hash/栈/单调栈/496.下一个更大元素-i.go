/*
 * @lc app=leetcode.cn id=496 lang=golang
 * @lcpr version=21909
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
// 右边更大更小，倒着来
package main

func nextGreaterElement(nums1, nums2 []int) []int {
	mp := map[int]int{}
	stack := []int{-1}
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 1 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		mp[num] = stack[len(stack)-1]
		stack = append(stack, num)
	}
	res := make([]int, len(nums1))
	for i, num := range nums1 {
		res[i] = mp[num]
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [4,1,2]\n[1,3,4,2].\n
// @lcpr case=end

// @lcpr case=start
// [2,4]\n[1,2,3,4].\n
// @lcpr case=end

*/
