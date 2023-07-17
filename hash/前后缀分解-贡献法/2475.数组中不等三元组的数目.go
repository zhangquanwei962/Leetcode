/*
 * @lc app=leetcode.cn id=2475 lang=golang
 * @lcpr version=21909
 *
 * [2475] 数组中不等三元组的数目
 */

// @lc code=start
package main

func unequalTriplets(nums []int) int {
	count := map[int]int{}
	for _, x := range nums {
		count[x]++
	}
	res, n, t := 0, len(nums), 0
	for _, v := range count {
		res, t = res+t*v*(n-t-v), t+v
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [4,4,2,4,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,1,1]\n
// @lcpr case=end

*/
