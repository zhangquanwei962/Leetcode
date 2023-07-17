/*
 * @lc app=leetcode.cn id=2491 lang=golang
 * @lcpr version=21909
 *
 * [2491] 划分技能点相等的团队
 */

// @lc code=start
// 典型哈希计数
package main

func dividePlayers(skill []int) (ans int64) {
	total := 0
	cnt := make(map[int]int)
	for _, x := range skill {
		total += x
		cnt[x]++
	}

	m := len(skill) / 2
	if total%m > 0 {
		return -1
	}

	s := total / m
	for x, c := range cnt {
		if c != cnt[s-x] {
			return -1
		}
		ans += int64(c * x * (s - x))
	}
	return ans / 2
}

// @lc code=end

/*
// @lcpr case=start
// [3,2,5,1,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [3,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,2,3]\n
// @lcpr case=end

*/
