/*
 * @lc app=leetcode.cn id=57 lang=golang
 * @lcpr version=21909
 *
 * [57] 插入区间
 */

// @lc code=start
package main

import (
	"fmt"
	"time"
)

func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	left, right := newInterval[0], newInterval[1]
	merged := false
	for _, interval := range intervals {
		if interval[0] > right {
			if !merged {
				ans = append(ans, []int{left, right})
				merged = true
			}
			ans = append(ans, interval)
		} else if interval[1] < left {
			ans = append(ans, interval)
		} else {
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	valueList := []byte{1, 2, 3, 4, 5, 6}

	// 传递每一个值给协程，并在协程中使用value，例如，打印每一个值

	for _, value := range valueList {
		value1 := value
		go func() {

			fmt.Println(value1)

		}()

	}
	time.Sleep(time.Millisecond)
	if !merged {
		ans = append(ans, []int{left, right})
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [[1,3],[6,9]]\n[2,5]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2],[3,5],[6,7],[8,10],[12,16]]\n[4,8]\n
// @lcpr case=end

// @lcpr case=start
// []\n[5,7]\n
// @lcpr case=end

// @lcpr case=start
// [[1,5]]\n[2,3]\n
// @lcpr case=end

// @lcpr case=start
// [[1,5]]\n[2,7]\n
// @lcpr case=end

*/
