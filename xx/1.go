package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	f := make([]int, n)
	for i := 0; i < n; i++ {
		left, right := -1, len(f)
		for left+1 < right {
			mid := left + (right-left)/2
			if intervals[i][0] >= intervals[mid][0] {
				left = mid
			} else {
				right = mid
			}
		}
		fmt.Println(left, right)
		j := right - 1

		f[i] = max(f[i], f[j]) + 1
	}
	return n - max(f...)
}

func max(a ...int) int {
	ma := a[0]
	for _, x := range a[1:] {
		if x > ma {
			ma = x
		}
	}
	return ma
}

func main() {
	a := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	x := eraseOverlapIntervals(a)
	fmt.Println(x)
}
