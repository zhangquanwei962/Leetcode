/*
 * @lc app=leetcode.cn id=23 lang=golang
 * @lcpr version=21909
 *
 * [23] 合并 K 个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 这个是(0,len(a)-1)
package main

func mergeTwoLists(a, b *ListNode) *ListNode {
	if a == nil || b == nil {
		if a == nil {
			return b
		}
		return a
	}

	var head ListNode
	tail, aPtr, bPtr := &head, a, b

	for aPtr != nil && bPtr != nil {
		if aPtr.Val < bPtr.Val {
			tail.Next = aPtr
			aPtr = aPtr.Next
		} else {
			tail.Next = bPtr
			bPtr = bPtr.Next
		}
		tail = tail.Next
	}

	if aPtr != nil {
		tail.Next = aPtr
	} else {
		tail.Next = bPtr
	}

	return head.Next
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l >= r {
		return nil
	}

	mid := (l + r) >> 1

	return mergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r))
}

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

// @lc code=end

/*
// @lcpr case=start
// [[1,4,5],[1,3,4],[2,6]]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [[]]\n
// @lcpr case=end

*/
