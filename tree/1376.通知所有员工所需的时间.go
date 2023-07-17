/*
 * @lc app=leetcode.cn id=1376 lang=golang
 * @lcpr version=21908
 *
 * [1376] 通知所有员工所需的时间
 */

// @lc code=start

// func numOfMinutes(n, headID int, manager, informTime []int) (ans int) {
//     g := make([][]int, n)
//     for i, m := range manager {
//         if m >= 0 {
//             g[m] = append(g[m], i) // 建树
//         }
//     }
//     var dfs func(int, int)
//     dfs = func(x, pathSum int) {
//         pathSum += informTime[x] // 累加递归路径上的 informTime[x]
//         ans = max(ans, pathSum)  // 更新答案的最大值
//         for _, y := range g[x] { // 遍历 x 的儿子 y（如果没有儿子就不会进入循环）
//             dfs(y, pathSum) // 继续递归
//         }
//     }
//     dfs(headID, 0) // 从根节点 headID 开始递归
//     return
// }

// func max(a, b int) int { if a < b { return b }; return a }
package main

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	g := make([][]int, n)
	for i, m := range manager {
		if m >= 0 {
			g[m] = append(g[m], i)
		}
	}

	var dfs func(int) int

	dfs = func(x int) (maxPathSum int) {
		for _, y := range g[x] {
			maxPathSum = max(maxPathSum, dfs(y))
		}
		return maxPathSum + informTime[x]
	}
	return dfs(headID)
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
// 1\n0\n[-1]\n[0]\n
// @lcpr case=end

// @lcpr case=start
// 6\n2\n[2,2,-1,2,2,2]\n[0,0,1,0,0,0]\n
// @lcpr case=end

*/
