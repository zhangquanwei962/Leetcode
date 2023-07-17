/*
 * @lc app=leetcode.cn id=916 lang=golang
 * @lcpr version=21909
 *
 * [916] 单词子集
 */

// @lc code=start
// 合并成一个单词，即求出最大出现次数
package main

func wordSubsets(A []string, B []string) []string {
	bmax := count("")
	for _, b := range B {
		bCount := count(b)
		for i := 0; i < 26; i++ {
			bmax[i] = max(bmax[i], bCount[i])
		}
	}

	var ans []string
search:
	for _, a := range A {
		aCount := count(a)
		for i := 0; i < 26; i++ {
			if aCount[i] < bmax[i] {
				continue search
			}
		}
		ans = append(ans, a)
	}

	return ans
}

func count(S string) []int {
	ans := make([]int, 26)
	for _, c := range S {
		ans[c-'a']++
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// ["amazon","apple","facebook","google","leetcode"]\n["e","o"]\n
// @lcpr case=end

// @lcpr case=start
// ["amazon","apple","facebook","google","leetcode"]\n["l","e"]\n
// @lcpr case=end

// @lcpr case=start
// ["amazon","apple","facebook","google","leetcode"]\n["e","oo"]\n
// @lcpr case=end

// @lcpr case=start
// ["amazon","apple","facebook","google","leetcode"]\n["lo","eo"]\n
// @lcpr case=end

// @lcpr case=start
// ["amazon","apple","facebook","google","leetcode"]\n["ec","oc","ceo"]\n
// @lcpr case=end

*/
