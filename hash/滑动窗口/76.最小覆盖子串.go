/*
 * @lc app=leetcode.cn id=76 lang=golang
 * @lcpr version=21909
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
// 滑动窗口思想
package main

import "math"

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	// 统计 t 中每个字符出现的次数
	targetCount := make(map[rune]int)
	for i := 0; i < len(t); i++ {
		targetCount[rune(t[i])]++
	}

	// 初始化窗口的左右指针和最小子串的起始位置和长度
	left := 0
	start, length := 0, math.MaxInt32

	// 记录窗口中已经包含的 t 中字符的数量
	count := len(targetCount)

	// 移动右指针扩大窗口，直到包含了 t 的所有字符
	for right, x := range s {
		// 如果当前字符在 t 中存在，则减少需要匹配的字符数
		if _, ok := targetCount[x]; ok {
			targetCount[x]--
			if targetCount[x] == 0 {
				count--
			}
		}

		// 当窗口包含了 t 的所有字符时，移动左指针缩小窗口
		for count == 0 {
			// 更新最小子串的起始位置和长度
			if right-left+1 < length {
				start = left
				length = right - left + 1
			}

			// 如果左指针对应的字符在 t 中存在，则增加需要匹配的字符数
			if _, ok := targetCount[rune(s[left])]; ok {
				targetCount[rune(s[left])]++
				if targetCount[rune(s[left])] > 0 {
					count++
				}
			}
			left++ // 移动左指针
		}
	}

	// 如果最小子串的长度没有改变，则表示找不到符合条件的子串
	if length == math.MaxInt32 {
		return ""
	}

	return s[start : start+length] // 返回最小子串
}

// func minWindow(s string, t string) string {
// 	m, n := len(s), len(t)
// 	if n > m {
// 		return ""
// 	}

// 	tMap := make(map[byte]int)
// 	for i := 0; i < n; i++ {
// 		tMap[t[i]]++
// 	}

// 	left, count := 0, 0
// 	minLen, minLeft := m+1, 0

// 	for right, ch := range s {
// 		if tMap[byte(ch)] > 0 {
// 			count++
// 		}
// 		tMap[byte(ch)]--
// 		for count == n {
// 			if right-left+1 < minLen {
// 				minLen = right - left + 1
// 				minLeft = left
// 			}
// 			if tMap[s[left]] == 0 {
// 				count--
// 				delete(tMap, s[left])
// 			}
// 			tMap[s[left]]++
// 			left++
// 		}
// 	}

// 	if minLen == m+1 {
// 		return ""
// 	}
// 	return s[minLeft : minLeft+minLen]
// }

// @lc code=end

/*
// @lcpr case=start
// "ADOBECODEBANC"\n"ABC"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n"a"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n"aa"\n
// @lcpr case=end

*/
