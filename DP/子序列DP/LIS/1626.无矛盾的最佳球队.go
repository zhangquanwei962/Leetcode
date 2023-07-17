/*
 * @lc app=leetcode.cn id=1626 lang=golang
 * @lcpr version=21909
 *
 * [1626] 无矛盾的最佳球队
 */

// @lc code=start
// 基于值域计算
package main

import "sort"

func bestTeamScore(scores, ages []int) int {
	n := len(scores)
	type pair struct{ score, age int }
	a := make([]pair, n)
	for i, age := range ages {
		a[i] = pair{scores[i], age}
	}
	sort.Slice(a, func(i, j int) bool {
		a, b := a[i], a[j]
		return a.score < b.score || a.score == b.score && a.age < b.age
	})

	maxSum := make([]int, max(ages)+1)
	for _, p := range a {
		maxSum[p.age] = max(maxSum[:p.age+1]) + p.score
	}
	return max(maxSum)
}

func max(a []int) int {
	mx := a[0]
	for _, x := range a {
		if x > mx {
			mx = x
		}
	}
	return mx
}

// func bestTeamScore(scores, ages []int) (ans int) {
//     n := len(scores)
//     type pair struct{ score, age int }
//     a := make([]pair, n)
//     for i, s := range scores {
//         a[i] = pair{s, ages[i]}
//     }
//     sort.Slice(a, func(i, j int) bool {
//         a, b := a[i], a[j]
//         return a.score < b.score || a.score == b.score && a.age < b.age
//     })

//     f := make([]int, n)
//     for i, p := range a {
//         for j, q := range a[:i] {
//             if q.age <= p.age {
//                 f[i] = max(f[i], f[j])
//             }
//         }
//         f[i] += p.score
//         ans = max(ans, f[i])
//     }
//     return
// }

// func max(a, b int) int { if a < b { return b }; return a }

// @lc code=end

/*
// @lcpr case=start
// [1,3,5,10,15]\n[1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [4,5,6,5]\n[2,1,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,5]\n[8,9,10,1]\n
// @lcpr case=end

*/
