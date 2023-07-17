/*
 * @lc app=leetcode.cn id=143 lang=golang
 * @lcpr version=21909
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

import "fmt"

// 876. 链表的中间结点
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	fmt.Printf("%d", slow.Val)
	return slow
}

// 206. 反转链表
func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func reorderList(head *ListNode) {
	mid := middleNode(head)
	head2 := reverseList(mid)
	for head2.Next != nil {
		nxt := head.Next
		nxt2 := head2.Next
		head.Next = head2
		head2.Next = nxt
		head = nxt
		head2 = nxt2
	}
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

*/
