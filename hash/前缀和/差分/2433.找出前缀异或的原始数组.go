/*
 * @lc app=leetcode.cn id=2433 lang=golang
 * @lcpr version=21909
 *
 * [2433] 找出前缀异或的原始数组
 */

// @lc code=start
// 等于相邻两个异或
package main

func findArray(pref []int) []int {
	result := []int{pref[0]}
	for i := 0; i < len(pref)-1; i++ {
		result = append(result, pref[i]^pref[i+1])
	}
	return result
}

// 输出：[1 3 1 7 5]

// @lc code=end

/*
// @lcpr case=start
// [5,2,0,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [13]\n
// @lcpr case=end

*/
