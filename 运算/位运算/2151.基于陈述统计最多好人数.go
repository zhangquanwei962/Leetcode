/*
 * @lc app=leetcode.cn id=2151 lang=golang
 * @lcpr version=21909
 *
 * [2151] 基于陈述统计最多好人数
 */

// @lc code=start
// 1是好人，0是坏人
package main

func maximumGood(statements [][]int) (ans int) {
	n := len(statements)
next:
	for i := 1; i < 1<<n; i++ {
		cnt := 0 // i中好人数目
		for j, row := range statements {
			if i>>j&1 > 0 { // 枚举 i 中的好人 j
				for k, st := range row {
					if st < 2 && st != i>>k&1 { // 该陈述与实际情况矛盾
						continue next
					}
				}
				cnt++
			}
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [[2,1,2],[1,2,2],[2,0,2]]\n
// @lcpr case=end

// @lcpr case=start
// [[2,0],[0,2]]\n
// @lcpr case=end

*/
