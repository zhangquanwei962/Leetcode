/*
 * @lc app=leetcode.cn id=216 lang=golang
 * @lcpr version=21909
 *
 * [216] 组合总和 III
 */

// @lc code=start
package main

func combinationSum3(k, n int) (ans [][]int) {
	path := []int{}
	var dfs func(int, int)
	dfs = func(i, t int) {
		d := k - len(path)              // 还要选 d 个数
		if t < 0 || t > (i*2-d+1)*d/2 { // 剪枝
			return
		}
		if d == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for j := i; j >= d; j-- {
			path = append(path, j)
			dfs(j-1, t-j)
			path = path[:len(path)-1]
		}
	}
	dfs(9, n)
	return
}

// @lc code=end

/*
// @lcpr case=start
// 3\n7\n
// @lcpr case=end

// @lcpr case=start
// 3\n9\n
// @lcpr case=end

// @lcpr case=start
// 4\n1\n
// @lcpr case=end

*/
