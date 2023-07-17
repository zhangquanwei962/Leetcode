/*
 * @lc app=leetcode.cn id=6898 lang=golang
 * @lcpr version=21909
 *
 * [6898] 字符串连接删减字母
 */

// @lc code=start
package main

func minimizeConcatenatedLength(words []string) int {
	memo := make([][][]int, len(words))
	for i := range memo {
		memo[i] = make([][]int, 26)
		for j := range memo[i] {
			memo[i][j] = make([]int, 26)
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}
	var dfs func(int, byte, byte) int
	dfs = func(i int, j, k byte) (res int) {
		if i == len(words) {
			return
		}

		p := &memo[i][j-'a'][k-'a']
		if *p != -1 {
			return *p
		}

		defer func() { *p = res }()

		w := words[i]

		res1 := dfs(i+1, j-'a', w[len(w)-1]-'a') - (w[0])

	}
}

// class Solution:
//     def minimizeConcatenatedLength(self, words: List[str]) -> int:
//         n = len(words)
//         @cache
//         def dfs(i,pre,last):
//             if i == n:
//                 return 0
//             # 放前面
//             if words[i][-1] == pre:
//                 res1 = dfs(i+1,words[i][0],last) + len(words[i]) -1
//             else:
//                 res1 = dfs(i+1,words[i][0],last) + len(words[i])
//             # 放后面
//             if words[i][0] == last:
//                 res2 = dfs(i+1,pre,words[i][-1]) + len(words[i]) -1
//             else:
//                 res2 = dfs(i+1,pre,words[i][-1]) + len(words[i])
//             return min(res1,res2)
//         return dfs(1,words[0][0],words[0][-1]) + len(words[0])

// @lc code=end

/*
// @lcpr case=start
// ["aa","ab","bc"]\n
// @lcpr case=end

// @lcpr case=start
// ["ab","b"]\n
// @lcpr case=end

// @lcpr case=start
// ["aaa","c","aba"]\n
// @lcpr case=end

*/
