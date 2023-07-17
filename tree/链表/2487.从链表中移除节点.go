/*
 * @lc app=leetcode.cn id=2487 lang=golang
 * @lcpr version=21909
 *
 * [2487] 从链表中移除节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//递归本质就是在倒着遍历链表
package main

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

func removeNodes(head *ListNode) *ListNode {
	head = reverseList(head)
	cur := head
	for cur.Next != nil {
		if cur.Val > cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return reverseList(head)
}

// func removeNodes(head *ListNode) *ListNode {
// 	if head.Next == nil || head == nil {
// 		return head
// 	}

// 	node := removeNodes(head.Next)
// 	if node.Val > head.Val {
// 		return node //删除head
// 	}

// 	head.Next = node //不删除head
// 	return head

// }

// @lc code=end

/*
// @lcpr case=start
// [5,2,13,3,8]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,1]\n
// @lcpr case=end

*/
