/*
 * @lc app=leetcode.cn id=128 lang=golang
 * @lcpr version=21909
 *
 * [128] 最长连续序列
 */

// @lc code=start
package main

import "fmt"

func longestConsecutive(nums []int) (longNum int) {
	numSet := map[int]struct{}{}

	for _, x := range nums {
		numSet[x] = struct{}{}
	}

	for num := range numSet {
		// fmt.Println(num)
		if _, ok := numSet[num-1]; !ok {
			curNum := num
			curLength := 1
			fmt.Println(curNum)

			for _, ok := numSet[curNum+1]; ok; _, ok = numSet[curNum+1] {
				curNum++
				curLength++
			}
			if curLength > longNum {
				longNum = curLength
			}
		}

	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [100,4,200,1,3,2]\n
// @lcpr case=end

// @lcpr case=start
// [0,3,7,2,5,8,4,6,0,1]\n
// @lcpr case=end

*/
