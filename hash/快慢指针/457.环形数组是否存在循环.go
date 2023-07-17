/*
 * @lc app=leetcode.cn id=457 lang=golang
 * @lcpr version=21909
 *
 * [457] 环形数组是否存在循环
 */

// @lc code=start
// O(n) O(1)
package main

func circularArrayLoop(nums []int) bool {
	n := len(nums)
	next := func(cur int) int {
		return ((cur+nums[cur])%n + n) % n // 保证返回值在 [0,n) 中
	}

	for i, x := range nums {
		if x == 0 {
			continue
		}
		slow, fast := i, next(i)
		// 判断非0且方向相同
		for nums[slow]*nums[fast] > 0 && nums[slow]*nums[next(fast)] > 0 {
			if slow == fast {
				// 排除i=next(i) 是自己本身
				if slow == next(slow) {
					break
				}
				return true
			}
			slow = next(slow)
			fast = next(next(fast))
		}
		// 每个点的访问情况，置0
		add := i
		for nums[add]*nums[next(add)] > 0 {
			tmp := add
			add = next(add)
			nums[tmp] = 0
		}
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// [2,-1,1,2,2]\n
// @lcpr case=end

// @lcpr case=start
// [-1,2]\n
// @lcpr case=end

// @lcpr case=start
// [-2,1,-1,-2,-2]\n
// @lcpr case=end

*/
