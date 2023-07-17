/*
 * @lc app=leetcode.cn id=2179 lang=golang
 * @lcpr version=21909
 *
 * [2179] 统计数组中好三元组数目
 */

// @lc code=start
// 置换
package main

func goodTriplets(nums1, nums2 []int) (ans int64) {
	n := len(nums1)
	p := make([]int, n)
	for i, v := range nums1 {
		p[v] = i
	}
	tree := make([]int, n+1)
	for i := 1; i < n-1; i++ {
		for j := p[nums2[i-1]] + 1; j <= n; j += j & -j { // 将 p[nums2[i-1]]+1 加入树状数组
			tree[j]++
		}
		y, less := p[nums2[i]], 0
		for j := y; j > 0; j &= j - 1 { // 计算 less
			less += tree[j]
		}
		ans += int64(less) * int64(n-1-y-(i-less))
	}
	return
}

// @lc code=end

/*
// @lcpr case=start
// [2,0,1,3]\n[0,1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [4,0,1,3,2]\n[4,1,0,2,3]\n
// @lcpr case=end

*/
