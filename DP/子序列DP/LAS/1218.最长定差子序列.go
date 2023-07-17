/*
 * @lc app=leetcode.cn id=1218 lang=golang
 * @lcpr version=21909
 *
 * [1218] 最长定差子序列
 */

// @lc code=start
package main

func longestSubsequence(arr []int, difference int) (ans int) {
	n := len(arr)
	f := make(map[int]int, n)
	for _, x := range arr {
		f[x] = f[x-difference] + 1
		if ans < f[x] {
			ans = f[x]
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,3,5,7]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,5,7,8,5,3,4,2,1]\n-2\n
// @lcpr case=end

*/
