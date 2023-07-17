/*
 * @lc app=leetcode.cn id=207 lang=golang
 * @lcpr version=21909
 *
 * [207] 课程表
 */

// @lc code=start
// O(n+m) O(m+n)
package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	result := []int{}

	for _, e := range prerequisites {
		x, y := e[0], e[1]
		g[y] = append(g[y], x)
		indeg[x]++
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		tmp := q
		q = nil
		for _, p := range tmp {
			result = append(result, p)
			for _, v := range g[p] {
				indeg[v]--
				if indeg[v] == 0 {
					q = append(q, v)
				}
			}

		}
	}
	return len(result) == numCourses
}

// @lc code=end

/*
// @lcpr case=start
// 2\n[[1,0]]\n
// @lcpr case=end

// @lcpr case=start
// 2\n[[1,0],[0,1]]\n
// @lcpr case=end

*/
