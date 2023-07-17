/*
 * @lc app=leetcode.cn id=503 lang=golang
 * @lcpr version=21909
 *
 * [503] 下一个更大元素 II
 */

// @lc code=start
package main

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	stk := []int{-1}
	ans := make([]int, n)
	for i := n*2 - 2; i >= 0; i-- {
		for len(stk) > 1 && nums[stk[len(stk)-1]] <= nums[i%n] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 1 {
			ans[i%n] = stk[len(stk)-1]
		} else {
			ans[i%n] = nums[stk[len(stk)-1]%n]
		}

		stk = append(stk, i%n)
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,3]\n
// @lcpr case=end

*/
