/*
 * @lc app=leetcode.cn id=206 lang=golang
 * @lcpr version=21909
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 狠狠理解递归
package main

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	first := head
// 	second := head.Next
// 	first.Next = reverseList(second.Next)
// 	second.Next = first
// 	return second

// }

// func reverseList(head *ListNode) *ListNode {
// 	var pre, cur *ListNode = nil, head
// 	for cur != nil {
// 		nxt := cur.Next
// 		cur.Next = pre
// 		pre = cur
// 		cur = nxt
// 	}
// 	return pre
// }

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
