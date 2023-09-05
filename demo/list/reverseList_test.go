package main

import "testing"

func generateListNode() *ListNode {
	head := &ListNode{val: 1}
	curr := head
	for i := 2; i <= 5; i++ {
		curr.Next = &ListNode{val: i}
		curr = curr.Next
	}
	return head
}

func TestMain1(t *testing.T) {
	head := &ListNode{val: 1}
	curr := head
	curr.Next = &ListNode{val: 2}
	curr = curr.Next
	for i := 2; i <= 5; i++ {
		curr.Next = &ListNode{val: i}
		curr = curr.Next
	}
	printList(head)
	// 1.链表反转
	//reverse(head)

	// 2.链表去重
	removeDuplicate(head)
	printList(head)

	// 3.链表反转去重
	//printList(reverseAndRemoveDuplicate(head))

	// 4.指定区间链表反转
	//printList(reverseBetween(head, 2, 4))

	// 5.删除链表的节点
	//printList(deleteNode(head, 1))

	// 6.K 个一组翻转链表
	printList(reverseKGroup(head, 3))
}

func ExampleReverseList3() {
	head := generateListNode()
	printList(head)
	printList(reverseList3(head))

	// output:
	// 1 -> 2 -> 3 -> 4 -> 5
	// 5 -> 4 -> 3 -> 2 -> 1
}

func ExampleReverseList2() {
	head := generateListNode()
	printList(head)
	printList(reverseList2(head))

	// output:
	// 1 -> 2 -> 3 -> 4 -> 5
	// 5 -> 4 -> 3 -> 2 -> 1
}

func ExampleReverseList() {
	head := generateListNode()
	printList(head)
	printList(reverseList(head))

	// output:
	// 1 -> 2 -> 3 -> 4 -> 5
	// 5 -> 4 -> 3 -> 2 -> 1
}

func ExampleReverseList4() {
	head := generateListNode()
	printList(head)
	printList(reverseList4(head))

	// output:
	// 1 -> 2 -> 3 -> 4 -> 5
	// 5 -> 4 -> 3 -> 2 -> 1
}

func ExampleReverseList5() {
	head := generateListNode()
	printList(head)
	printList(reverseList5(head))

	// output:
	// 1 -> 2 -> 3 -> 4 -> 5
	// 5 -> 4 -> 3 -> 2 -> 1
}

// 2.链表去重
func ExampleRemoveDuplicate() {
	head := &ListNode{val: 1}
	curr := head
	curr.Next = &ListNode{val: 2}
	curr = curr.Next
	for i := 2; i <= 5; i++ {
		curr.Next = &ListNode{val: i}
		curr = curr.Next
	}
	printList(head)
	printList(removeDuplicate(head))

	// output:
	// 1 -> 2 -> 2 -> 3 -> 4 -> 5
	// 1 -> 2 -> 3 -> 4 -> 5
}

// 3.链表反转去重
func ExampleReverseAndRemoveDuplicate() {
	head := &ListNode{val: 1}
	curr := head
	curr.Next = &ListNode{val: 2}
	curr = curr.Next
	for i := 2; i <= 5; i++ {
		curr.Next = &ListNode{val: i}
		curr = curr.Next
	}
	printList(head)
	printList(reverseAndRemoveDuplicate(head))

	// output:
	// 1 -> 2 -> 2 -> 3 -> 4 -> 5
	// 5 -> 4 -> 3 -> 2 -> 1
}

// 4.指定区间链表反转
func ExampleReverseBetween() {
	head := generateListNode()
	printList(head)
	printList(reverseBetween(head, 2, 4))

	// output:
	// 1 -> 2 -> 3 -> 4 -> 5
	// 1 -> 4 -> 3 -> 2 -> 5
}

// 5.删除链表的节点
func ExampleDeleteNode() {
	head := generateListNode()
	printList(head)
	printList(deleteNode(head, 2))

	// output:
	// 1 -> 2 -> 3 -> 4 -> 5
	// 1 -> 3 -> 4 -> 5
}

// 6.K 个一组翻转链表
func ExampleReverseKGroup() {
	head := generateListNode()
	printList(head)
	printList(reverseKGroup(head, 2))

	// output:
	// 1 -> 2 -> 3 -> 4 -> 5
	// 2 -> 1 -> 4 -> 3 -> 5
}

// 6.K 个一组翻转链表
func ExampleReverseKGroup2() {
	head := generateListNode()
	printList(head)
	printList(reverseKGroup(head, 3))

	// output:
	// 1 -> 2 -> 3 -> 4 -> 5
	// 3 -> 2 -> 1 -> 4 -> 5
}

// 7.两两交换链表中的节点
func ExampleSswapPairs() {
	head := generateListNode()
	printList(head)
	printList(swapPairs(head))

	// output:
	// 1 -> 2 -> 3 -> 4 -> 5
	// 2 -> 1 -> 4 -> 3 -> 5
}
