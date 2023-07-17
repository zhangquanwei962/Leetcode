/*
 * @lc app=leetcode.cn id=2670 lang=golang
 * @lcpr version=21909
 *
 * [2670] 找出不同元素数目差数组
 */

// @lc code=start
// 模拟，先算后缀，一遍算一遍遍历
// O(n) O(n)
package main

func distinctDifferenceArray(nums []int) []int {
	n := len(nums)
	suf := make([]int, n+1)
	set := make(map[int]struct{})

	for i := n - 1; i >= 0; i-- {
		set[nums[i]] = struct{}{}
		suf[i] = len(set)
	}

	set = make(map[int]struct{})
	ans := make([]int, n)
	for i, x := range nums {
		set[x] = struct{}{}
		ans[i] = len(set) - suf[i+1]
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,3,4,2]\n
// @lcpr case=end

*/
