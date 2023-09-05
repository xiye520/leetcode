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

// temp.next = node2
// node1.next = node2.next
// node2.next = node1
// 最后令temp=node1
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

// 递归解法
// 如果链表中至少有两个节点，则在两两交换链表中的节点之后，原始链表的头节点变成新的链表的第二个节点，原始链表的第二个节点变成新的链表的头节点。
// 链表中的其余节点的两两交换可以递归地实现。在对链表中的其余节点递归地两两交换之后，更新节点之间的指针关系，即可完成整个链表的两两交换。
//
// 用 head 表示原始链表的头节点，新的链表的第二个节点，用 newHead 表示新的链表的头节点，原始链表的第二个节点，
// 则原始链表中的其余节点的头节点是 newHead.next。令 head.next = swapPairs(newHead.next)，表示将其余节点进行两两交换，
// 交换后的新的头节点为 head 的下一个节点。然后令 newHead.next = head，即完成了所有节点的交换。最后返回新的链表的头节点 newHead。

func swapPairs3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}
