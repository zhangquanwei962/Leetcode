/*
 * @lc app=leetcode.cn id=1170 lang=golang
 * @lcpr version=21909
 *
 * [1170] 比较字符串最小字母出现频次
 */

// @lc code=start
package main

func f(s string) int {
	cnt, ch := 0, 'z'

	for _, x := range s {
		if x < ch {
			cnt, ch = 1, x
		} else if ch == x {
			cnt++
		}
	}
	return cnt
}

func numSmallerByFrequency(queries []string, words []string) []int {
	count := make([]int, 12)
	for _, s := range words {
		count[f(s)]++
	}

	for i := 9; i >= 1; i-- {
		count[i] += count[i+1]
	}

	res := make([]int, len(queries))
	for i, s := range queries {
		res[i] = count[f(s)+1]
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// ["cbd"]\n["zaaaz"]\n
// @lcpr case=end

// @lcpr case=start
// ["bbb","cc"]\n["a","aa","aaa","aaaa"]\n
// @lcpr case=end

*/
