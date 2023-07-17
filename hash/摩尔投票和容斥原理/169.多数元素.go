/*
 * @lc app=leetcode.cn id=169 lang=golang
 * @lcpr version=21908
 *
 * [169] 多数元素
 */

// @lc code=start

// 摩尔投票
package main

func majorityElement(nums []int) int {
	people, cnt := -1, 0
	for _, x := range nums {
		if x == people {
			cnt++
		} else if cnt--; cnt < 0 {
			people = x
			cnt = 1
		}
	}
	return people
}

// @lc code=end

/*
// @lcpr case=start
// [3,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [2,2,1,1,1,2,2]\n
// @lcpr case=end

*/
