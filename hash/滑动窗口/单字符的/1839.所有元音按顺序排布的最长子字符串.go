/*
 * @lc app=leetcode.cn id=1839 lang=golang
 * @lcpr version=21908
 *
 * [1839] 所有元音按顺序排布的最长子字符串
 */

// @lc code=start
package main

func longestBeautifulSubstring(word string) (ans int) {
	cnt := make([]int, 'v') // 声明一个大小为 'z'+1 的 int 数组
	length := 0

	for right := 0; right < len(word); right++ {
		cnt[word[right]]++
		length++
		if right > 0 && word[right] < word[right-1] {
			cnt = make([]int, 'v') // 使用 make 函数重新分配内存来清空 cnt 数组
			cnt[word[right]]++
			length = 1
		}
		if cnt['a'] > 0 && cnt['e'] > 0 && cnt['i'] > 0 && cnt['o'] > 0 && cnt['u'] > 0 {
			ans = max(ans, length)
		}
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
// "aeiaaioaaaaeiiiiouuuooaauuaeiu"\n
// @lcpr case=end

// @lcpr case=start
// "aeeeiiiioooauuuaeiou"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n
// @lcpr case=end

*/
