/*
 * @lc app=leetcode.cn id=1887 lang=golang
 * @lcpr version=21909
 *
 * [1887] 使数组元素相等的减少操作次数
 */

// @lc code=start
// 相当于求出每个数在数组中是第几小的
// 读题理解题意
package main

import "sort"

func reductionOperations(nums []int) (ans int) {
	sort.Ints(nums)
	for i, k, n := 0, 0, len(nums); k < n; k++ {
		start := i
		for ; i < n && nums[i] == nums[start]; i++ {
		}
		ans += (i - start) * k // 区间内[start,i)都是第k小
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [5,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,2,2,3]\n
// @lcpr case=end

*/
