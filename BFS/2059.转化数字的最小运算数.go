/*
 * @lc app=leetcode.cn id=2059 lang=golang
 * @lcpr version=21908
 *
 * [2059] 转化数字的最小运算数
 */

// @lc code=start
package main

func minimumOperations(nums []int, start int, goal int) int {
	vis := [1001]bool{}
	vis[start] = true
	q := []int{start}

	for step := 1; q != nil; step++ {
		tmp := q
		q = nil
		for _, v := range tmp {
			for _, num := range nums {
				for _, x := range []int{v + num, v - num, v ^ num} {
					if x == goal {
						return step
					}

					if 0 <= x && x <= 1000 && !vis[x] {
						vis[x] = true
						q = append(q, x)
					}
				}
			}
		}
	}
	return -1
}

// @lc code=end

/*
// @lcpr case=start
// [2,4,12]\n2\n12\n
// @lcpr case=end

// @lcpr case=start
// [3,5,7]\n0\n-4\n
// @lcpr case=end

// @lcpr case=start
// [2,8,16]\n0\n1\n
// @lcpr case=end

*/
