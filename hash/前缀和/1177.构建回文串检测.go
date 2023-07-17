/*
 * @lc app=leetcode.cn id=1177 lang=golang
 * @lcpr version=21909
 *
 * [1177] 构建回文串检测
 */

// @lc code=start
// 看到quary就想到前缀和，然后就是回文数字说白了就是偶数个
// 最多有一个奇数个的，想到用异或（1的个数）
package main

import "math/bits"

func canMakePaliQueries(s string, queries [][]int) []bool {
	n := len(s)
	count := make([]int, n+1)
	for i := 0; i < n; i++ {
		count[i+1] = count[i] ^ (1 << (s[i] - 'a'))
	}
	res := make([]bool, len(queries))
	for i, query := range queries {
		l := query[0]
		r := query[1]
		k := query[2]
		x := count[r+1] ^ count[l]
		// for x > 0 {
		// 	x &= x - 1
		// 	bits++
		// }
		bits := bits.OnesCount(uint(x))
		res[i] = bits <= k*2+1
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// "abcda"\n[[3,3,0],[1,2,0],[0,3,1],[0,3,2],[0,4,1]]\n
// @lcpr case=end

*/
