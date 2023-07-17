/*
 * @lc app=leetcode.cn id=77 lang=golang
 * @lcpr version=21909
 *
 * [77] 组合
 */

// @lc code=start
package main

func combine(n int, k int) (ans [][]int) {
	path := []int{}
	var dfs func(int)
	dfs = func(i int) {
		d := k - len(path)
		if d == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		for j := i; j >= d; j-- {
			path = append(path, j)
			dfs(j - 1)
			path = path[:len(path)-1]
		}
	}

	dfs(n)
	return
}

// func combine(n, k int) (ans [][]int) {
//     path := []int{}
//     var dfs func(int)
//     dfs = func(i int) {
//         d := k - len(path) // 还要选 d 个数
//         if d == 0 {
//             ans = append(ans, append([]int(nil), path...))
//             return
//         }
//         // 不选 i
//         if i > d {
//             dfs(i - 1)
//         }
//         // 选 i
//         path = append(path, i)
//         dfs(i - 1)
//         path = path[:len(path)-1]
//     }
//     dfs(n)
//     return
// }

// @lc code=end

/*
// @lcpr case=start
// 4\n2\n
// @lcpr case=end

// @lcpr case=start
// 1\n1\n
// @lcpr case=end

*/
