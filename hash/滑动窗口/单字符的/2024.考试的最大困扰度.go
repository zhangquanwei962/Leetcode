/*
 * @lc app=leetcode.cn id=2024 lang=golang
 * @lcpr version=21908
 *
 * [2024] 考试的最大困扰度
 */

// @lc code=start
package main

func maxConsecutiveAnswers(answerKey string, k int) (ans int) {
	cnt := make([]int, 26)
	left := 0

	for right, ch := range answerKey {
		cnt[ch-'A']++
		for cnt['T'-'A'] > k && cnt['F'-'A'] > k {
			cnt[answerKey[left]-'A']--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
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
// "TTFF"\n2\n
// @lcpr case=end

// @lcpr case=start
// "TFFT"\n1\n
// @lcpr case=end

// @lcpr case=start
// "TTFTTFTT"\n1\n
// @lcpr case=end

*/
