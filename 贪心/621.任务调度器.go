/*
 * @lc app=leetcode.cn id=621 lang=golang
 * @lcpr version=21909
 *
 * [621] 任务调度器
 */

// @lc code=start
// 考虑执行次数最多的任务
package main

func leastInterval(tasks []byte, n int) int {
	cnt := map[byte]int{}
	for _, t := range tasks {
		cnt[t]++
	}

	maxExec, maxExecCnt := 0, 0
	for _, c := range cnt {
		if c > maxExec {
			maxExec, maxExecCnt = c, 1
		} else if c == maxExec {
			maxExecCnt++
		}
	}

	if time := (maxExec-1)*(n+1) + maxExecCnt; time > len(tasks) {
		return time
	}
	return len(tasks)
}

// @lc code=end

/*
// @lcpr case=start
// ["A","A","A","B","B","B"]\n2\n
// @lcpr case=end

// @lcpr case=start
// ["A","A","A","B","B","B"]\n0\n
// @lcpr case=end

// @lcpr case=start
// ["A","A","A","A","A","A","B","C","D","E","F","G"]\n2\n
// @lcpr case=end

*/
