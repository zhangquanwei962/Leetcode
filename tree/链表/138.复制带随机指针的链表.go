/*
 * @lc app=leetcode.cn id=138 lang=golang
 * @lcpr version=21909
 *
 * [138] 复制带随机指针的链表
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
package main

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	cur := head
	nodeMap := make(map[*Node]*Node)

	for cur != nil {
		nodeMap[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		nodeMap[cur].Next = nodeMap[cur.Next]
		nodeMap[cur].Random = nodeMap[cur.Random]
		cur = cur.Next
	}

	return nodeMap[head]
}

// var cachedNode map[*Node]*Node

// func deepCopy(node *Node) *Node {
// 	if node == nil {
// 		return node
// 	}

// 	if n, has := cachedNode[node]; has {
// 		return n
// 	}

// 	newNode := &Node{Val: node.Val}
// 	cachedNode[node] = newNode
// 	newNode.Next = deepCopy(node.Next)
// 	newNode.Random = deepCopy(node.Random)
// 	return newNode
// }
// func copyRandomList(head *Node) *Node {
// 	cachedNode = map[*Node]*Node{}
// 	return deepCopy(head)
// }

// @lc code=end

/*
// @lcpr case=start
// [[7,null],[13,0],[11,4],[10,2],[1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,1],[2,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[3,null],[3,0],[3,null]]\n
// @lcpr case=end

*/
