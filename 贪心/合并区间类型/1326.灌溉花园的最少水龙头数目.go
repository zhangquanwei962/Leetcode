/*
 * @lc app=leetcode.cn id=1326 lang=golang
 * @lcpr version=21909
 *
 * [1326] 灌溉花园的最少水龙头数目
 */

// @lc code=start
// 如果到 n-1 时没有返回 -1，那么肯定可以到 n。
package main

func minTaps(n int, ranges []int) (ans int) {
	rightMost := make([]int, n+1)
	for i, r := range ranges {
		left := max(i-r, 0)
		rightMost[left] = max(rightMost[left], i+r)
	}

	curRight := 0                     // 已建造的桥的右端点
	nextRight := 0                    // 下一座桥的右端点的最大值
	for i, r := range rightMost[:n] { // 注意这里没有遍历到 n，因为它已经是终点了
		nextRight = max(nextRight, r)
		if i == curRight { // 到达已建造的桥的右端点
			if i == nextRight { // 无论怎么造桥，都无法从 i 到 i+1
				return -1
			}
			curRight = nextRight // 造一座桥
			ans++
		}
	}
	return
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
// 5\n[3,4,1,1,0,0]\n
// @lcpr case=end

// @lcpr case=start
// 3\n[0,0,0,0]\n
// @lcpr case=end

*/
