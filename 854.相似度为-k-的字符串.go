/*
 * @lc app=leetcode.cn id=854 lang=golang
 * @lcpr version=21909
 *
 * [854] 相似度为 K 的字符串
 */

// @lc code=start
// 数据范围小，可以枚举,找最短路
/*
1、首先BFS过程中存下每一种中间状态，因为第一次搜到肯定是到达的最小步数，重复搜索直接跳过；2、进行字符串之间的状态转移的时候，如果某一位已经和最终状态相同了，那么这位就不要动它，因为动了只会使得步数增加，不是最优解；3、同样的，每次交换，至少要是使得中间状态的某一位变成最终状态的字符，否则这一步就白交换了，得到的也一定不是最小次数；4、每一步只考虑交换一位，并且要寻找所有的这一位
后边可以交的位置拓展出一些新的中间状态
2.进行字符串之间的状态转移的时候，如果某一位已经和最终状态相同了，那么这位就不要动它，因为动了只会使得步数增加，不是最优解；
3、同样的，每次交换，至少要是使得中间状态的某一位变成最终状态的字符，
否则这一步就白交换了，得到的也一定不是最小次数
*/
package main

func kSimilarity(s1, s2 string) int {
	var s, t []byte
	for i := range s1 {
		if s1[i] != s2[i] {
			s = append(s, s1[i])
			t = append(t, s2[i])
		}
	}
	n := len(s)
	if n == 0 {
		return 0
	}

	minSwap := func(i int) int {
		diff := 0
		for j := i; j < n; j++ {
			if s[j] != t[j] {
				diff++
			}
		}
		return (diff + 1) / 2
	}

	ans := n - 1
	vis := make(map[string]bool, n)
	var dfs func(int, int, string)
	dfs = func(i, cost int, s string) {
		if cost > ans {
			return
		}
		if vis[s] == true {
			return
		}
		for i < n && s[i] == t[i] {
			i++
		}
		if i == n {
			ans = min(ans, cost)
			return
		}
		// 当前状态的交换次数下限大于等于当前的最小交换次数
		if cost+minSwap(i) >= ans {
			return
		}
		vis[s] = true
		s1 := []byte(s)
		for j := i + 1; j < n; j++ {
			if s[j] == t[i] {
				s1[i], s1[j] = s1[j], s1[i]
				dfs(i+1, cost+1, string(s1))
				s1[i], s1[j] = s1[j], s1[i]
			}
		}
	}
	dfs(0, 0, string(s))
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// package main

// func kSimilarity(s1 string, s2 string) (step int) {
// 	type pair struct {
// 		s string
// 		i int
// 	}

// 	q := []pair{{s1, 0}}
// 	vis := map[string]bool{s1: true}
// 	for n := len(s1); ; step++ {
// 		tmp := q
// 		q = nil
// 		for _, p := range tmp {
// 			s, i := p.s, p.i
// 			if s == s2 {
// 				return
// 			}
// 			for i < n && s[i] == s2[i] {
// 				i++
// 			}
// 			t := []byte(s)
// 			for j := i + 1; j < n; j++ {
// 				if s[j] == s2[i] && s[j] != s2[j] { // 剪枝，只在 s[j] != s2[j] 时去交换
// 					t[i], t[j] = t[j], t[i]
// 					if t := string(t); !vis[t] {
// 						vis[t] = true
// 						q = append(q, pair{t, i + 1})
// 					}
// 					t[i], t[j] = t[j], t[i]
// 				}
// 			}
// 		}
// 	}
// }

// @lc code=end

/*
// @lcpr case=start
// "ab"\n"ba"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n"bca"\n
// @lcpr case=end

*/
