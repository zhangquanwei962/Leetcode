/*
 * @lc app=leetcode.cn id=904 lang=golang
 * @lcpr version=21909
 *
 * [904] 水果成篮
 */

// @lc code=start
package main

func totalFruit(fruits []int) (ans int) {
	left := 0
	cnt := make(map[int]int)
	for right, x := range fruits {
		cnt[x]++
		for len(cnt) > 2 {
			cnt[fruits[left]]--
			if cnt[fruits[left]] == 0 {
				delete(cnt, fruits[left])
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
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
// [1,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,2,2]\n
// @lcpr case=end

// @lcpr case=start
// [3,3,3,1,2,1,1,2,3,3,4]\n
// @lcpr case=end

*/
