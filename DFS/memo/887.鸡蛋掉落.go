/*
 * @lc app=leetcode.cn id=887 lang=golang
 * @lcpr version=21909
 *
 * [887] 鸡蛋掉落
 */

// @lc code=start
package main

func superEggDrop(k int, n int) int {
	ans := 1
	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, n+2)
		for j := range f[i] {
			f[i][j] = -1
		}
	}
	// i 蛋的个数 j 机会
	var caclF func(int, int) int
	caclF = func(i, j int) (res int) {
		if j == 1 || i == 1 {
			return j + 1
		}
		p := &f[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		res = caclF(i-1, j-1) + caclF(i, j-1)
		return res
	}
	for caclF(k, ans) < n+1 {
		ans++
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// 1\n2\n
// @lcpr case=end

// @lcpr case=start
// 2\n6\n
// @lcpr case=end

// @lcpr case=start
// 3\n14\n
// @lcpr case=end

*/
