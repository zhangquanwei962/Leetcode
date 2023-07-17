/*
 * @lc app=leetcode.cn id=1071 lang=golang
 * @lcpr version=21908
 *
 * [1071] 字符串的最大公因子
 */

// @lc code=start

/*
* 该实现的时间复杂度为 O(L + log N)，
* 其中 L 是 str1 和 str2 的长度之和，
* log N 是计算最大公约数的时间复杂度。空间复杂度为 O(1)。
 */
package main

import "strings"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcd1(a, b int) int {
	if b != 0 {
		return gcd1(b, a%b)
	}
	return a
}

func gcdOfStrings(str1 string, str2 string) string {
	L1, L2 := len(str1), len(str2)
	g := gcd(L1, L2)
	if str1[:g] != str2[:g] {
		return ""
	}

	x := str1[:g]

	if strings.Repeat(x, L1/g) != str1 || strings.Repeat(x, L2/g) != str2 {
		return ""
	}
	return x
}

// @lc code=end

/*
// @lcpr case=start
// "ABCABC"\n"ABC"\n
// @lcpr case=end

// @lcpr case=start
// "ABABAB"\n"ABAB"\n
// @lcpr case=end

// @lcpr case=start
// "LEET"\n"CODE"\n
// @lcpr case=end

*/
