/*
 * @lc app=leetcode.cn id=424 lang=golang
 * @lcpr version=21908
 *
 * [424] 替换后的最长重复字符
 */

// @lc code=start
package main

func characterReplacement(s string, k int) (maxSize int) {
	cnt := ['Z' + 1]int{}
	left, maxCnt := 0, 0
	for right, ch := range s {
		cnt[ch]++
		maxCnt = max(maxCnt, cnt[ch])
		for right-left+1-maxCnt > k {
			cnt[s[left]]--
			left++
		}
		maxSize = max(maxSize, right-left+1)
	}
	return
}

// func characterReplacement(s string, k int) int {
//     cnt := [26]int{}
//     left, maxCnt, maxSize := 0, 0, 0
//     for right, ch := range s {
//         cnt[ch-'A']++
//         maxCnt = max(maxCnt, cnt[ch-'A'])
//         if right-left+1-maxCnt > k {
//             cnt[s[left]-'A']--
//             left++
//         }
//         maxSize = max(maxSize, right-left+1)
//     }
//     return maxSize
// }

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

/*
// @lcpr case=start
// "ABAB"\n2\n
// @lcpr case=end

// @lcpr case=start
// "AABABBA"\n1\n
// @lcpr case=end

*/
