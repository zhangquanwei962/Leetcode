/*
 * @lc app=leetcode.cn id=435 lang=golang
 * @lcpr version=21909
 *
 * [435] 无重叠区间
 */

// @lc code=start
// 可以重叠，找下一个大于intervals[i][0]
// 正难则fan
package main

import "sort"

func eraseOverlapIntervals(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	maxRight := points[0][1]
	ans := 1
	for _, p := range points {
		if p[0] >= maxRight {
			maxRight = p[1]
			ans++
		}
	}
	return len(points) - ans
}

// func eraseOverlapIntervals(intervals [][]int) int {
// 	sort.Slice(intervals, func(i, j int) bool {
// 		a, b := intervals[i], intervals[j]
// 		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
// 	})

// 	n := len(intervals)
// 	f := []int{}

// 	for _, e := range intervals {
// 		// [0,n)
// 		j := sort.Search(len(f), func(k int) bool {
// 			return f[k] > e[0]
// 		})
// 		if j == len(f) {
// 			f = append(f, e[1])
// 		} else if f[j] > e[1] {
// 			f[j] = e[1]
// 		}

// 	}

// 	return n - len(f)
// }

/*
class Solution:
    def eraseOverlapIntervals(self, intervals: List[List[int]]) -> int:
        intervals.sort(key=lambda x: x[0])
        f = [0] * len(intervals)

        for i in range(len(intervals)):
            for j in range(i):
                if intervals[i][0] >= intervals[j][1]:
                    f[i] = max(f[i], f[j])
            f[i] += 1
        return len(intervals) - max(f)
*/

// func eraseOverlapIntervals(intervals [][]int) int {
// 	n := len(intervals)
// 	if n == 0 {
// 		return 0
// 	}
// 	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
// 	ans, right := 1, intervals[0][1]
// 	for _, p := range intervals[1:] {
// 		if p[0] >= right {
// 			ans++
// 			right = p[1]
// 		}
// 	}
// 	return n - ans
// }

/*
这段代码是一个用于消除重叠区间的函数 `eraseOverlapIntervals`，它接受一个二维整数切片 `intervals` 作为输入，返回消除重叠后剩余区间的最小数量。

我们来逐步解释代码的工作原理：

1. 第一行导入了必要的包 `sort`。

2. 使用 `sort.Slice` 函数对 `intervals` 进行排序。排序规则是，首先按照区间的起始值进行升序排序，如果起始值相同，则按照区间的结束值进行降序排序。这样可以确保较早结束的区间排在前面，方便后续处理。

3. 初始化变量 `n`，表示区间的总数，以及切片 `f`，用于存储不重叠的区间的结束值。

4. 使用 `range` 循环遍历已经排序的 `intervals` 中的每个区间。

5. 在每次循环中，使用 `sort.Search` 函数在 `f` 中查找第一个大于当前区间起始值的索引，并赋值给变量 `j`。这里利用了二分查找的性质。

6. 根据 `j` 的值进行不同的处理：
   - 若 `j` 等于 `f` 的长度，说明当前区间与之前的所有区间都无重叠，将当前区间的结束值添加到 `f` 中。
   - 否则，更新 `f[j]` 为当前区间的结束值，这样可以保证 `f` 中存储的结束值是当前起始值下的最小结束值。

7. 最后，函数返回 `n - len(f)`，即剩余不重叠区间的数量。

通过以上步骤，该函数可以在时间复杂度为 O(nlogn) 的情况下，有效地消除重叠区间。希望这能帮助到您，如果有任何进一步的问题，请随时提问。
*/

// @lc code=end

/*
// @lcpr case=start
// [[1,2],[2,3],[3,4],[1,3]]\n
// @lcpr case=end

// @lcpr case=start
// [ [1,2], [1,2], [1,2] ]\n
// @lcpr case=end

// @lcpr case=start
// [ [1,2], [2,3] ]\n
// @lcpr case=end

*/
