/*
 * @lc app=leetcode.cn id=278 lang=golang
 * @lcpr version=21908
 *
 * [278] 第一个错误的版本
 */

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
package main

func firstBadVersion(n int) int {
	// return sort.Search(n, func(i int) bool {
	// 	return isBadVersion(i)
	// })
	var binary func(int, int) int
	binary = func(i1, i2 int) int {
		for i1+1 < i2 {
			mid := i1 + (i2-i1)/2
			if isBadVersion(mid) {
				i2 = mid
			} else {
				i1 = mid
			}
		}
		return i2
	}
	return binary(-1, n+1)
}

// @lc code=end

/*
// @lcpr case=start
// 5\n4\n
// @lcpr case=end

// @lcpr case=start
// 1\n1\n
// @lcpr case=end

*/
