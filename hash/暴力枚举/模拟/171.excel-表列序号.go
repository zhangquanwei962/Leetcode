/*
 * @lc app=leetcode.cn id=171 lang=golang
 * @lcpr version=21909
 *
 * [171] Excel 表列序号
 */

// @lc code=start
package main

func titleToNumber(columnTitle string) int {
	ans := 0
	for i := 0; i < len(columnTitle); i++ {
		ans = ans*26 + int(columnTitle[i]-'A'+1)
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// "A"\n
// @lcpr case=end

// @lcpr case=start
// "AB"\n
// @lcpr case=end

// @lcpr case=start
// "ZY"\n
// @lcpr case=end

*/
