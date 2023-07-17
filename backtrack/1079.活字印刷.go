/*
 * @lc app=leetcode.cn id=1079 lang=golang
 * @lcpr version=21907
 *
 * [1079] 活字印刷
 */

// @lc code=start
package main

func numTilePossibilities(tiles string) int {
	count := make(map[rune]int)
	for _, t := range tiles {
		count[t]++
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if i == 0 {
			return 1
		}

		res := 1 // 也是一个解
		for t := range count {
			if count[t] > 0 {
				count[t]--
				res += dfs(i - 1)
				count[t]++
			}
		}
		return res
	}
	return dfs(len(tiles)) - 1

}

// @lc code=end

/*
// @lcpr case=start
// "AAB"\n
// @lcpr case=end

// @lcpr case=start
// "AAABBC"\n
// @lcpr case=end

// @lcpr case=start
// "V"\n
// @lcpr case=end

*/
