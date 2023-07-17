/*
 * @lc app=leetcode.cn id=1171 lang=golang
 * @lcpr version=21909
 *
 * [1171] 从链表中删去总和值为零的连续节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 哈希表还是比较巧妙的
// 有正有负，那么第二次出现，中间的就是0
// 其实就是hash前缀和
package main

func removeZeroSumSublists(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	seen := map[int]*ListNode{}

	prefix := 0

	for node := dummy; node != nil; node = node.Next {
		prefix += node.Val
		seen[prefix] = node
	}

	prefix = 0
	for node := dummy; node != nil; node = node.Next {
		prefix += node.Val
		node.Next = seen[prefix].Next
	}

	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,-3,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,-3,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,-3,-2]\n
// @lcpr case=end

*/
