/*
 * @lc app=leetcode.cn id=1029 lang=golang
 * @lcpr version=21909
 *
 * [1029] 两地调度
 */

// @lc code=start
package main

import "sort"

func twoCitySchedCost(costs [][]int) (totalCos int) {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][1]-costs[i][0] < costs[j][1]-costs[j][0]
	})

	n := len(costs) / 2
	for i := n; i < len(costs); i++ {
		totalCos += costs[i][0]
	}
	for i := 0; i < n; i++ {
		totalCos += costs[i][1]
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [[10,20],[30,200],[400,50],[30,20]]\n
// @lcpr case=end

// @lcpr case=start
// [[259,770],[448,54],[926,667],[184,139],[840,118],[577,469]]\n
// @lcpr case=end

// @lcpr case=start
// [[515,563],[451,713],[537,709],[343,819],[855,779],[457,60],[650,359],[631,42]]\n
// @lcpr case=end

*/
