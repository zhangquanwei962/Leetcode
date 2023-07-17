/*
 * @lc app=leetcode.cn id=1438 lang=golang
 * @lcpr version=21909
 *
 * [1438] 绝对差不超过限制的最长连续子数组
 */

// @lc code=start
// 哈希表维护窗口内的数字
package main

func longestSubarray(nums []int, limit int) (ans int) {
	cnt := make(map[int]int)
	left := 0
	for right, x := range nums {
		cnt[x]++
		for {
			mx, mn := x, x
			for k := range cnt {
				mx = max(mx, k)
				mn = min(mn, k)
			}
			if mx-mn <= limit {
				break
			}
			y := nums[left]
			if cnt[y]--; cnt[y] == 0 {
				delete(cnt, y)
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [8,2,4,7]\n4\n
// @lcpr case=end

// @lcpr case=start
// [10,1,2,4,7,2]\n5\n
// @lcpr case=end

// @lcpr case=start
// [4,2,2,2,4,4,2,2]\n0\n
// @lcpr case=end

*/
