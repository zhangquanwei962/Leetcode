/*
 * @lc app=leetcode.cn id=165 lang=golang
 * @lcpr version=21909
 *
 * [165] 比较版本号
 */

// @lc code=start
// 两个序列，就是双指针
package main

func compareVersion(version1 string, version2 string) int {
	m, n := len(version1), len(version2)
	i, j := 0, 0

	for i < m || j < n {
		x := 0
		for ; i < m && version1[i] != '.'; i++ {
			x = 10*x + int(version1[i]-'0')
		}
		i++
		y := 0
		for ; j < n && version2[j] != '.'; j++ {
			y = 10*y + int(version2[j]-'0')
		}

		j++

		if x > y {
			return 1
		}

		if x < y {
			return -1
		}

	}
	return 0
}

// @lc code=end

/*
// @lcpr case=start
// "1.01"\n"1.001"\n
// @lcpr case=end

// @lcpr case=start
// "1.0"\n"1.0.0"\n
// @lcpr case=end

// @lcpr case=start
// "0.1"\n"1.1"\n
// @lcpr case=end

*/
