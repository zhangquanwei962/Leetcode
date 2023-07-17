/*
 * @lc app=leetcode.cn id=524 lang=golang
 * @lcpr version=21909
 *
 * [524] 通过删除字母匹配到字典里最长单词
 */

// @lc code=start
// 两个子序列，有单调性，双指针呼之欲出
// O(d*(m+n)) O(1)
package main

func findLongestWord(s string, dictionary []string) (ans string) {
	for _, t := range dictionary {
		i := 0
		for j := range s {
			if s[j] == t[i] {
				i++
				if i == len(t) {
					if len(t) > len(ans) || len(t) == len(ans) && t < ans {
						ans = t
					}
					break
				}
			}
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// "abpcplea"\n["ale","apple","monkey","plea"]\n
// @lcpr case=end

// @lcpr case=start
// "abpcplea"\n["a","b","c"]\n
// @lcpr case=end

*/
