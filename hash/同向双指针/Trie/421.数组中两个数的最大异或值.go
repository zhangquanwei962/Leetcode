/*
 * @lc app=leetcode.cn id=421 lang=golang
 * @lcpr version=21909
 *
 * [421] 数组中两个数的最大异或值
 */

// @lc code=start
// O(nlogC) O(nlogC)
package main

const highBit int = 30

type Trie struct {
	left, right *Trie
}

func (t *Trie) add(num int) {
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.left == nil {
				cur.left = &Trie{}
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = &Trie{}
			}
			cur = cur.right
		}
	}
}

func (t *Trie) check(num int) (x int) {
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			// a_i 的第 k 个二进制位为 0，应当往表示 1 的子节点 right 走
			if cur.right != nil {
				cur = cur.right
				x = x*2 + 1
			} else {
				cur = cur.left
				x = x * 2
			}
		} else {
			// a_i 的第 k 个二进制位为 1，应当往表示 0 的子节点 left 走
			if cur.left != nil {
				cur = cur.left
				x = x*2 + 1
			} else {
				cur = cur.right
				x = x * 2
			}
		}
	}
	return
}

func findMaximumXOR(nums []int) (x int) {
	root := &Trie{}
	for i := 1; i < len(nums); i++ {
		// 将 nums[i-1] 放入字典树，此时 nums[0 .. i-1] 都在字典树中
		root.add(nums[i-1])
		// 将 nums[i] 看作 ai，找出最大的 x 更新答案
		x = max(x, root.check(nums[i]))
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

/*
// @lcpr case=start
// [3,10,5,25,2,8]\n
// @lcpr case=end

// @lcpr case=start
// [14,70,53,83,49,91,36,80,92,51,66,70]\n
// @lcpr case=end

*/
