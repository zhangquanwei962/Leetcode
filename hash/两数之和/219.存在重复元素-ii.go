/*
 * @lc app=leetcode.cn id=219 lang=golang
 * @lcpr version=21909
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
package main

func containsNearbyDuplicate(nums []int, k int) bool {
	pos := map[int]int{}
	for i, num := range nums {
		if p, ok := pos[num]; ok && i-p <= k {
			return true
		}
		pos[num] = i
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,0,1,1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,1,2,3]\n2\n
// @lcpr case=end

*/
