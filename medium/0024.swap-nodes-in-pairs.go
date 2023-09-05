package medium

/** 24. 两两交换链表中的节点
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

示例 1：
输入：head = [1,2,3,4]
输出：[2,1,4,3]

示例 2：
输入：head = []
输出：[]

示例 3：
输入：head = [1]
输出：[1]

https://leetcode.cn/problems/swap-nodes-in-pairs/description/
*/

func swapPairs(head *ListNode) *ListNode {
	var dummy ListNode
	dummy.Next = head
	curr := &dummy
	for head.Next != nil && head.Next.Next != nil {
		node1 := curr.Next
		node2 := curr.Next.Next

		curr.Next = node2
		node1.Next = node2.Next
		node2.Next = node1

		curr = node1
	}
	return dummy.Next
}

func swapPairs2(head *ListNode) *ListNode {
	var dummy ListNode
	dummy.Next = head
	curr := &dummy
	for curr.Next != nil && curr.Next.Next != nil {
		node1 := curr.Next
		node2 := curr.Next.Next

		// 当前的下一个指向2
		curr.Next = node2
		// node1的下一个节点更新位指向2的下一个
		node1.Next = node2.Next
		// 2指向1，即反转
		node2.Next = node1
		// 更新当前节点为1（1和2已经反转了）
		curr = node1
	}

	return dummy.Next
}
