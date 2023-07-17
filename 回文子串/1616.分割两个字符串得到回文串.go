/*
 * @lc app=leetcode.cn id=1616 lang=golang
 * @lcpr version=21909
 *
 * [1616] 分割两个字符串得到回文串
 */

// @lc code=start
// 相向双指针求出前后最长匹配
// 判断a或者b中间部分是不是回文串
package main

// 如果从 s[i] 到 s[j] 是回文串则返回 true，否则返回 false
func isPalindrome(s string, i, j int) bool {
	for i < j && s[i] == s[j] {
		i++
		j--
	}
	return i >= j
}

// 如果 a_prefix + b_suffix 可以构成回文串则返回 true，否则返回 false
func check(a, b string) bool {
	i, j := 0, len(a)-1         // 相向双指针
	for i < j && a[i] == b[j] { // 前后缀尽量匹配
		i++
		j--
	}
	return isPalindrome(a, i, j) || isPalindrome(b, i, j)
}

func checkPalindromeFormation(a, b string) bool {
	return check(a, b) || check(b, a)
}

// @lc code=end

/*
// @lcpr case=start
// "x"\n"y"\n
// @lcpr case=end

// @lcpr case=start
// "abdef"\n"fecab"\n
// @lcpr case=end

// @lcpr case=start
// "ulacfd"\n"jizalu"\n
// @lcpr case=end

*/
