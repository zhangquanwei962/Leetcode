/*
 * @lc app=leetcode.cn id=1690 lang=golang
 * @lcpr version=21909
 *
 * [1690] 石子游戏 VII
 */

// @lc code=start
// https://leetcode.cn/problems/stone-game-vii/solution/by-liupengsay-tnly/
package main

func stoneGameVII(stones []int) int {
	n := len(stones)
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + stones[i] // 前缀和
	}
	// f[i][j]剩余stones[i:j+1]时先手能取得的最大分数
	// 能定义成最大差值吗？？？（应该也是最大差值）
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = max(pre[j+1]-pre[i+1]-f[i+1][j],
				pre[j]-pre[i]-f[i][j-1])
		}
	}
	return f[0][n-1]
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
// [5,3,1,4,2]\n
// @lcpr case=end

// @lcpr case=start
// [7,90,5,1,100,10,10,2]\n
// @lcpr case=end

*/
