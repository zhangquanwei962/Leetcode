/*
 * @lc app=leetcode.cn id=373 lang=golang
 * @lcpr version=21908
 *
 * [373] 查找和最小的 K 对数字
 */

// @lc code=start
// (i,j)出堆时，(i,j+1)入堆
// (i,0)都要入堆
package main

import "container/heap"

// for i := range h {
// h[i] = tuple{nums1[i] + nums2[0], i, 0}
// }
func kSmallestPairs(nums1, nums2 []int, k int) [][]int {
	n, m := len(nums1), len(nums2)
	ans := make([][]int, 0, min(k, n*m)) // 预分配空间
	h := hp{{nums1[0] + nums2[0], 0, 0}}

	for len(h) > 0 && len(ans) < k {
		p := heap.Pop(&h).(tuple)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j == 0 && i+1 < n {
			heap.Push(&h, tuple{nums1[i+1] + nums2[0], i + 1, 0})
		}
		if j+1 < m {
			heap.Push(&h, tuple{nums1[i] + nums2[j+1], i, j + 1})
		}
	}
	return ans
}

type tuple struct{ s, i, j int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].s < h[j].s }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// @lc code=end

/*
// @lcpr case=start
// [1,7,11]\n[2,4,6]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,1,2]\n[1,2,3]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n[3]\n3\n
// @lcpr case=end

*/
