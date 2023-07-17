/*
 * @lc app=leetcode.cn id=1850 lang=golang
 * @lcpr version=21909
 *
 * [1850] 邻位交换的最小次数
 */

// @lc code=start
// 31 下一个排列
// 枚举查找不同的，贪心选择最小的j，不断变换
package main

func getMinSwaps(num string, k int) int {
	var num_k []byte
	num_k = []byte(num)
	for i := 0; i < k; i++ {
		nextPermutation(num_k)
	}

	n := len(num)
	ans := 0
	num1 := []byte(num)

	for i := 0; i < n; i++ {
		if num1[i] != num_k[i] {
			for j := i + 1; j < n; j++ {
				if num1[j] == num_k[i] {
					for k := j - 1; k >= i; k-- {
						ans++
						num1[k], num1[k+1] = num1[k+1], num1[k]
					}
					break
				}
			}
		}
	}
	return ans
}

func nextPermutation(nums []byte) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	reverse(nums[i+1:])
}

func reverse(a []byte) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

// @lc code=end

/*
// @lcpr case=start
// "5489355142"\n4\n
// @lcpr case=end

// @lcpr case=start
// "11112"\n4\n
// @lcpr case=end

// @lcpr case=start
// "00123"\n1\n
// @lcpr case=end

*/
