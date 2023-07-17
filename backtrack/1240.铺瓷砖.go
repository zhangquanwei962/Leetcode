/*
 * @lc app=leetcode.cn id=1240 lang=golang
 * @lcpr version=21909
 *
 * [1240] 铺瓷砖
 */

// @lc code=start
package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func tilingRectangle(n int, m int) int {
	ans := max(n, m)
	memo := make([][]bool, n)
	for i := range memo {
		memo[i] = make([]bool, m)
	}

	isAvaibalbe := func(x, y, k int) bool {
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				if memo[x+i][y+j] {
					return false
				}
			}
		}
		return true
	}

	fillUp := func(x, y, k int, val bool) {
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				memo[x+i][y+j] = val
			}
		}
	}

	var dfs func(int, int, int)
	dfs = func(x, y, cnt int) {
		if cnt >= ans { //剪枝
			return
		}

		if x >= n { //递归边界
			ans = cnt
			return
		}

		if y >= m { //搜索下一行
			dfs(x+1, 0, cnt)
			return
		}

		if memo[x][y] { //如果当前已经被覆盖，尝试下一个位置
			dfs(x, y+1, cnt)
			return
		}

		for k := min(n-x, m-y); k >= 1 && isAvaibalbe(x, y, k); k-- {
			// 将长度为 k 的正方形区域标记覆盖
			fillUp(x, y, k, true)
			// 跳过 k 个位置开始检测
			dfs(x, y+k, cnt+1)
			fillUp(x, y, k, false)
		}

	}
	dfs(0, 0, 0)
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// 2\n3\n
// @lcpr case=end

// @lcpr case=start
// 5\n8\n
// @lcpr case=end

// @lcpr case=start
// 11\n13\n
// @lcpr case=end

*/
