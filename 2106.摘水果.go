/*
 * @lc app=leetcode.cn id=2106 lang=golang
 * @lcpr version=21907
 *
 * [2106] 摘水果
 */

// @lc code=start
package main

import "sort"

func maxTotalFruits(fruits [][]int, startPos, k int) int {
	n := len(fruits)
	// 向左最远能到 fruits[left][0]
	left := sort.Search(n, func(i int) bool { return fruits[i][0] >= startPos-k })
	right, s := left, 0
	for ; right < n && fruits[right][0] <= startPos; right++ {
		s += fruits[right][1] // 从 fruits[left][0] 到 startPos 的水果数
	}
	ans := s
	for ; right < n && fruits[right][0] <= startPos+k; right++ {
		s += fruits[right][1] // 枚举最右位置为 fruits[right][0]
		for fruits[right][0]*2-fruits[left][0]-startPos > k &&
			fruits[right][0]-fruits[left][0]*2+startPos > k {
			s -= fruits[left][1] // fruits[left][0] 无法到达
			left++
		}
		ans = max(ans, s) // 更新答案最大值
	}
	return ans
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
// [[2,8],[6,3],[8,6]]\n5\n4\n
// @lcpr case=end

// @lcpr case=start
// [[0,9],[4,1],[5,7],[6,2],[7,4],[10,9]]\n5\n4\n
// @lcpr case=end

// @lcpr case=start
// [[0,3],[6,4],[8,5]]\n3\n2\n
// @lcpr case=end

*/
