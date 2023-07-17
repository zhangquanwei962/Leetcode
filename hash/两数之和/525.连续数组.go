/*
 * @lc app=leetcode.cn id=525 lang=golang
 * @lcpr version=21909
 *
 * [525] 连续数组
 */

// @lc code=start
package main

func findMaxLength(nums []int) int {
	mp := map[int]int{0: -1}
	cur, ans := 0, 0

	for i, x := range nums {
		if x == 0 {
			cur--
		} else {
			cur++
		}

		if _, ok := mp[cur]; ok {
			ans = max(ans, i-mp[cur])
		} else {
			mp[cur] = i
		}
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
// [0,1]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,0]\n
// @lcpr case=end

*/
