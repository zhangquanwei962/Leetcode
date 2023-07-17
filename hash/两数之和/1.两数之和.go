/*
 * @lc app=leetcode.cn id=1 lang=golang
 * @lcpr version=21908
 *
 * [1] 两数之和
 */

// @lc code=start
// 先查询，再更新
package main

func twoSum(nums []int, target int) []int {
	cnt := make(map[int]int)

	for i, x := range nums {
		if j, ok := cnt[target-x]; ok {
			return []int{i, j}
		}
		cnt[x] = i

	}
	return nil
}

// @lc code=end

/*
// @lcpr case=start
// [2,7,11,15]\n9\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [3,3]\n6\n
// @lcpr case=end

*/
