/*
 * @lc app=leetcode.cn id=1299 lang=golang
 * @lcpr version=21909
 *
 * [1299] 将每个元素替换为右侧最大元素
 */

// @lc code=start
package main

func replaceElements(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)
	ans[n-1] = -1
	for i := n - 2; i >= 0; i-- {
		ans[i] = max(ans[i+1], arr[i+1])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [17,18,5,4,6,1]\n
// @lcpr case=end

// @lcpr case=start
// [400]\n
// @lcpr case=end

*/
