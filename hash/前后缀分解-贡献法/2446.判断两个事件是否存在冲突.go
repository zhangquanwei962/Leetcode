/*
 * @lc app=leetcode.cn id=2446 lang=golang
 * @lcpr version=21907
 *
 * [2446] 判断两个事件是否存在冲突
 */

// @lc code=start
package main

func haveConflict(event1 []string, event2 []string) bool {
	return event1[1] >= event2[0] && event1[0] <= event2[1]
}

// @lc code=end

/*
// @lcpr case=start
// ["01:15","02:00"]\n["02:00","03:00"]\n
// @lcpr case=end

// @lcpr case=start
// ["01:00","02:00"]\n["01:20","03:00"]\n
// @lcpr case=end

// @lcpr case=start
// ["10:00","11:00"]\n["14:00","15:00"]\n
// @lcpr case=end

*/
