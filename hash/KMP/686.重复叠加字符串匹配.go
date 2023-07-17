/*
 * @lc app=leetcode.cn id=686 lang=golang
 * @lcpr version=21909
 *
 * [686] 重复叠加字符串匹配
 */

// @lc code=start
// 要知道这么一个道理，能成为字串肯定从哪个位置开始都能匹配上
package main

func repeatedStringMatch(a string, b string) int {
	m, n := len(a), len(b)
	if n == 0 {
		return 0
	}
	Str := func(a, b string) int {
		ne := make([]int, n)
		ne[0] = -1
		for i, j := 1, -1; i < n; i++ {
			for j > -1 && b[i] != b[j+1] {
				j = ne[j]
			}
			if b[i] == b[j+1] {
				j++
			}
			ne[i] = j
		}

		for i, j := 0, -1; i-j <= m; i++ {
			for j > -1 && a[i%m] != b[j+1] {
				j = ne[j]
			}
			if a[i%m] == b[j+1] {
				j++
			}
			if j == n-1 {
				return i - n + 1
			}
		}
		return -1
	}
	index := Str(a, b)
	if index == -1 {
		return -1
	}
	if m-index >= n {
		return 1
	}

	return (n+index-m-1)/m + 2

}

// @lc code=end

/*
// @lcpr case=start
// "abcd"\n"cdabcdab"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n"aa"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n"a"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n"wxyz"\n
// @lcpr case=end

*/
