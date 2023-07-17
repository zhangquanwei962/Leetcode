/*
 * @lc app=leetcode.cn id=1250 lang=golang
 * @lcpr version=21909
 *
 * [1250] 检查「好数组」
 */

// @lc code=start
// 两个数，ax+by=1,如果a b互质，那么一定有解，推论到n个；
// 而a b互质要求gcd等于1
// O(n+logm) O(1)
package main

func isGoodArray(nums []int) bool {
	g := 0
	for _, x := range nums {
		g = gcd(x, g)
	}
	return g == 1
}

// func gcd(a, b int) int {
// 	if b == 0 {
// 		return a
// 	}
// 	return gcd(b, a%b)
// }

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [12,5,7,23]\n
// @lcpr case=end

// @lcpr case=start
// [29,6,10]\n
// @lcpr case=end

// @lcpr case=start
// [3,6]\n
// @lcpr case=end

*/
