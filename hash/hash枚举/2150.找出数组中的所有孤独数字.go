/*
 * @lc app=leetcode.cn id=2150 lang=golang
 * @lcpr version=21909
 *
 * [2150] 找出数组中的所有孤独数字
 */

// @lc code=start
package main

func findLonely(nums []int) (ans []int) {
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
	}
	for v, c := range cnt {
		if c == 1 && cnt[v+1] == 0 && cnt[v-1] == 0 {
			ans = append(ans, v)
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [10,6,5,8]\n
// @lcpr case=end

// @lcpr case=start
// [1,3,5,3]\n
// @lcpr case=end

*/
