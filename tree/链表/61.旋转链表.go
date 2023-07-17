/*
 * @lc app=leetcode.cn id=61 lang=golang
 * @lcpr version=21909
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 断链为环
package main

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	n := 1
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		n++
	}

	add := n - k%n
	if add == n {
		return head
	}
	// 成环
	iter.Next = head
	for add > 0 {
		iter = iter.Next
		add--
	}
	// 断开
	ret := iter.Next
	iter.Next = nil
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2]\n4\n
// @lcpr case=end

*/
