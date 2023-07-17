/*
 * @lc app=leetcode.cn id=2134 lang=golang
 * @lcpr version=21909
 *
 * [2134] 最少交换次数来组合所有的 1 II
 */

// @lc code=start
// 计算总共1的窗口最少能有多少个0
package main

func minSwaps(nums []int) int {
	tot1 := 0 // 1的总个数
	for _, num := range nums {
		tot1 += num
	}
	/*
		统计出数组中1的个数为cnt，那么固定一个cnt的窗口，
		里面0的个数最少有多少个，就是我们要交换的最少次数
		因为里面的1一定是不用动的，而里面有多少个0就对应外面
		有多少个0，而这里的数组是环形的，因此要遍历两遍
	*/
	cnt1, maxCnt1 := 0, 0
	nums = append(nums, nums...) // 断环成链,两倍
	for i, num := range nums {
		cnt1 += num    // 1的个数
		if i >= tot1 { // 滑窗
			cnt1 -= nums[i-tot1]
			if cnt1 > maxCnt1 {
				maxCnt1 = cnt1
			}
		}
	}
	return tot1 - maxCnt1
}

// @lc code=end

/*
// @lcpr case=start
// [0,1,0,1,1,0,0]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,1,1,0,0,1,1,0]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,0,0,1]\n
// @lcpr case=end

*/
