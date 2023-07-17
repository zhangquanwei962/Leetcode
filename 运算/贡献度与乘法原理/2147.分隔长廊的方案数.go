/*
 * @lc app=leetcode.cn id=2147 lang=golang
 * @lcpr version=21909
 *
 * [2147] 分隔长廊的方案数
 */

// @lc code=start
package main

func numberOfWays(corridor string) int {
	ans, cntS, pre := 1, 0, 0
	for i, ch := range corridor {
		if ch == 'S' {
			// 对第 3,5,7,... 个座位，可以在其到其左侧最近座位之间的任意一个位置放置屏风
			cntS++
			if cntS >= 3 && cntS&1 == 1 {
				ans = ans * (i - pre) % (1e9 + 7)
			}
			pre = i // 记录上一个座位的位置
		}
	}
	if cntS == 0 || cntS&1 == 1 { // 座位个数不能为 0 或奇数
		return 0
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// "SSPPSPS"\n
// @lcpr case=end

// @lcpr case=start
// "PPSPSP"\n
// @lcpr case=end

// @lcpr case=start
// "S"\n
// @lcpr case=end

*/
