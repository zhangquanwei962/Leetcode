/*
 * @lc app=leetcode.cn id=2514 lang=golang
 * @lcpr version=21909
 *
 * [2514] 统计同位异构字符串数目
 */

// @lc code=start
// 费马小定理
// 首先，每个单词自己是互相独立的，因此分别计算每个单词的同位异构字符串的数目，再用乘法原理相乘。
// b^-1(mod p)=pow(b,p-2)(mod p)
// 3!/2!  O(n+logM) O(siga)
package main

import "strings"

const mod int = 1e9 + 7

func countAnagrams(s string) int {
	ans, mul := 1, 1
	for _, s := range strings.Split(s, " ") {
		cnt := [26]int{}
		for i, c := range s {
			cnt[c-'a']++
			mul = mul * cnt[c-'a'] % mod
			ans = ans * (i + 1) % mod
		}
	}
	return ans * pow(mul, mod-2) % mod
}

func pow(x, n int) int {
	res := 1
	for ; n != 0; n >>= 1 {
		if n&1 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

// func pow(x, n int) int {
// 	res := 1
// 	for ; n > 0; n >>= 1 {
// 		if n&1 > 0 {
// 			res = res * x % mod
// 		}
// 		x = x * x % mod
// 	}
// 	return res
// }

// @lc code=end

/*
// @lcpr case=start
// "too hot"\n
// @lcpr case=end

// @lcpr case=start
// "aa"\n
// @lcpr case=end

*/
