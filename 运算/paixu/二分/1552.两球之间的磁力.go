/*
 * @lc app=leetcode.cn id=1552 lang=golang
 * @lcpr version=21908
 *
 * [1552] 两球之间的磁力
 */

// @lc code=start
// 最小化最大值,是<= 返回right
// 最大化最小值，是>= 返回left
package main

import "sort"

func maxDistance(position []int, m int) int {
	sort.Ints(position)

	var check func(int) bool
	check = func(d int) bool {
		cnt, pre := 1, position[0]
		for _, p := range position {
			if p >= pre+d {
				cnt++
				pre = p
			}

		}
		return cnt < m

	}
	var binary func(int, int) int
	binary = func(left, right int) int {
		for left+1 < right {
			mid := left + (right-left)/2
			if check(mid) {
				right = mid
			} else {
				left = mid
			}
		}
		return left
	}

	return binary(0, (position[len(position)-1]-position[0])/(m-1)+1)

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,7]\n3\n
// @lcpr case=end

// @lcpr case=start
// [5,4,3,2,1,1000000000]\n2\n
// @lcpr case=end

*/
