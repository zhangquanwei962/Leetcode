/*
 * @lc app=leetcode.cn id=148 lang=golang
 * @lcpr version=21909
 *
 * [148] 排序链表
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

func mergeList(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

func sort2(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	// 找到的是head -> nil的数量的中间，如果
	// 是len(list)+(nil)偶数是后面那个
	slow, fast := head, head
	for fast != tail && fast.Next != tail {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 等价于
	// for fast != tail {
	// 	slow = slow.Next
	// 	fast = fast.Next
	// 	if fast != tail {
	// 		fast = fast.Next
	// 	}
	// }

	mid := slow
	return mergeList(sort2(head, mid), sort2(mid, tail))
}

func sortList(head *ListNode) *ListNode {
	return sort2(head, nil)
}

// @lc code=end

/*
// @lcpr case=start
// [4,2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [-1,5,3,4,0]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
