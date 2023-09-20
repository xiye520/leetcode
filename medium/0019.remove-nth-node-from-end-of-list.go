package medium

/** 19. 删除链表的倒数第 N 个结点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

https://leetcode.cn/problems/remove-nth-node-from-end-of-list/?envType=study-plan-v2&envId=top-interview-150
*/

// 方法一：计算链表长度
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	curr := dummy
	for i := 0; i < length-n; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	return dummy.Next
}

func getLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

// 二、栈
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}

// 三、双指针
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	var dummy ListNode
	dummy.Next = head
	first := head
	second := &dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next

	return dummy.Next
}
