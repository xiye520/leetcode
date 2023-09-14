package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type ListNode struct {
	val  int
	Next *ListNode
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode //  前一个节点
	curr := head       //  当前节点
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// 1.链表反转
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	curr := head
	for curr != nil {
		prev, curr, curr.Next = curr, curr.Next, prev
		//next := curr.Next
		//curr.Next = prev
		//prev = curr
		//curr = next
	}
	return prev
}

func reverseList5(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	lst := reverseList5(head.Next)
	head.Next.Next = head
	head.Next = nil
	return lst
}

//思路：可以利用头插法的特点：插入顺序与最后形成的链表顺序相反
/*
	将链表中的每个位置的数据取取出来，新建一个表头
	使用头插法进行插入
	这样只需遍历一次原列表O(n)，并且空间复杂度仅为O(1)
*/
func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dummy *ListNode = nil
	var curr *ListNode = nil
	for head != nil {
		curr = head
		// 将 curr 从 head 中摘除
		head = head.Next
		// 将 curr 插入到 dummy 的头部
		curr.Next = dummy
		dummy = curr
	}
	return dummy
}

// 2.链表去重
func removeDuplicate(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var dummy ListNode
	prev := &dummy
	curr := head
	for curr != nil {
		val := curr.val
		if curr.Next != nil && val == curr.Next.val {
			curr = curr.Next
			continue
		}

		prev.Next = curr
		prev = prev.Next
		curr = curr.Next
	}
	return dummy.Next
}

// 3.链表反转去重
func reverseAndRemoveDuplicate(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	curr := head
	for curr != nil {
		if prev != nil && prev.val == curr.val {
			curr = curr.Next
			continue
		}
		prev, curr, curr.Next = curr, curr.Next, prev
		//next := curr.Next
		//curr.Next = prev
		//prev = curr
		//curr = next
	}
	return prev
}

func main() {
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
	//reverseAndRemoveDuplicate(head)

	// 4.指定区间链表反转
	//reverseBetween(head, 2, 4)

	// 5.删除链表的节点
	//deleteNode(head, 1)

	// 6.K 个一组翻转链表
	printList(reverseKGroup(head, 3))

}

func printList(node *ListNode) {
	var bf bytes.Buffer
	first := true
	for node != nil {
		if first {
			first = false
			bf.WriteString(strconv.Itoa(node.val))
		} else {
			bf.WriteString(fmt.Sprintf(" -> %d", node.val))
		}
		node = node.Next
	}
	fmt.Println(bf.String())
}

// 4.指定区间链表反转
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}

	var dummy ListNode
	dummy.Next = head
	start := &dummy
	for i := 1; i < m; i++ {
		start = start.Next
	}

	end := start.Next
	for i := m; i < n; i++ {
		end.Next, end.Next.Next, start.Next = end.Next.Next, start.Next, end.Next
	}

	return dummy.Next
}

// 5.删除链表的节点
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	var dummy ListNode
	prev := &dummy
	curr := head
	for curr != nil {
		if curr.val == val {
			curr = curr.Next
			// 更新上一节点指向的下一个指针
			prev.Next = curr
			continue
		}
		// 当前值设置给prev
		prev.Next = curr
		prev = prev.Next
		curr = curr.Next
	}
	return dummy.Next
}

// 单指针
func deleteNode1(head *ListNode, val int) *ListNode {
	if head.val == val {
		return head.Next
	}
	pre := head
	for pre.Next != nil && pre.Next.val != val {
		pre = pre.Next
	}
	// 上面循环结束后，pre的下一个元素就是要删除的节点
	if pre.Next != nil {
		pre.Next = pre.Next.Next
	}

	return head
}

// 双指针
func deleteNode2(head *ListNode, val int) *ListNode {
	if head.val == val {
		return head.Next
	}

	pre, cur := head, head
	for cur != nil && cur.val != val {
		pre, cur = cur, cur.Next
	}
	if cur.Next != nil {
		pre.Next = cur.Next
	}

	return head
}

// 递归
func deleteNode3(head *ListNode, val int) *ListNode {
	if head.val == val {
		return head.Next
	}

	head.Next = deleteNode3(head.Next, val)

	return head
}

// 6.K 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	arr := make([]int, 0, 10)
	for head != nil {
		arr = append(arr, head.val)
		head = head.Next
	}
	reverseKNumbers(arr, k)
	newHead := buildListNode(arr)
	return newHead
}

// 7.两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	var dummy ListNode
	dummy.Next = head
	curr := &dummy
	for curr.Next != nil && curr.Next.Next != nil {
		node1 := curr.Next
		node2 := curr.Next.Next

		curr.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		curr = node1
	}

	return dummy.Next
}

// 8. 重排链表(leetcode 143)
func reorderList(head *ListNode) *ListNode {
	nodes := make([]*ListNode, 0, 10)
	curr := head
	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}
	start, end := 0, len(nodes)-1
	for start < end {
		nodes[start].Next = nodes[end]
		start++
		nodes[end].Next = nodes[start]
		end--
	}
	nodes[start].Next = nil
	return head
}
