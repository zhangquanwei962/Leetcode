/*
 * @lc app=leetcode.cn id=2193 lang=golang
 * @lcpr version=21909
 *
 * [2193] 得到回文串的最少操作次数
 */

// @lc code=start
package main

func minMovesToMakePalindrome(s string) int {
	// 每个字母出现的次数
	freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	ans := 0
	// 前一半和后一半
	left := make(map[byte][]int)
	right := make(map[byte][]int)
	lcnt := 0
	rcnt := 0

	// 统计「组间交换」的操作次数
	for i := 0; i < len(s); i++ {
		c := s[i]
		if len(left[c])+1 <= freq[c]/2 {
			// 属于前一半
			lcnt++
			left[c] = append(left[c], lcnt)
			ans += (i - lcnt + 1)
		} else {
			// 属于后一半
			rcnt++
			right[c] = append(right[c], rcnt)
		}
	}

	// 如果长度为奇数，需要在前一半末尾添加一个中心字母
	if len(s)%2 == 1 {
		for c, occ := range freq {
			if occ%2 == 1 {
				lcnt++
				left[c] = append(left[c], lcnt)
				break
			}
		}
	}

	// 得到排列
	perm := make([]int, (len(s)+1)/2)
	for c, rlist := range right {
		llist := left[c]
		for i := 0; i < len(rlist); i++ {
			perm[rlist[len(rlist)-i-1]-1] = llist[i]
		}
	}
	reverse(perm)

	// 计算逆序对，统计「组内交换」的操作次数
	// 暴力法
	getBruteForce := func() int {
		n := len(perm)
		cnt := 0
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if perm[i] > perm[j] {
					cnt++
				}
			}
		}
		return cnt
	}

	return getBruteForce() + ans
}

func reverse(arr []int) {
	left := 0
	right := len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

// @lc code=end

/*
// @lcpr case=start
// "aabb"\n
// @lcpr case=end

// @lcpr case=start
// "letelt"\n
// @lcpr case=end

*/
