/*
 * @lc app=leetcode.cn id=2135 lang=golang
 * @lcpr version=21909
 *
 * [2135] 统计追加字母可以获得的单词数
 */

// @lc code=start
// 并不真的需要排序，而是直接用哈希表记录即可
// 逆向思维，增加等于减少，集合位运算判断
package main

func wordCount(startWords []string, targetWords []string) (ans int) {
	has := map[int]bool{}
	for _, word := range startWords {
		mask := 0
		for _, ch := range word {
			mask |= 1 << (ch - 'a')
		}
		has[mask] = true
	}

	for _, word := range targetWords {
		mask := 0
		for _, ch := range word {
			mask |= 1 << (ch - 'a')
		}
		for _, ch := range word {
			if has[mask^(1<<(ch-'a'))] { // 去掉这个字符
				ans++
				break
			}
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// ["ant","act","tack"]\n["tack","act","acti"]\n
// @lcpr case=end

// @lcpr case=start
// ["ab","a"]\n["abc","abcd"]\n
// @lcpr case=end

*/
