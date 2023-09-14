package medium

import (
	"bytes"
	"fmt"
	"strconv"
)

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

func generateListNode() *ListNode {
	head := &ListNode{Val: 1}
	curr := head
	for i := 2; i <= 5; i++ {
		curr.Next = &ListNode{Val: i}
		curr = curr.Next
	}
	return head
}

func printList(node *ListNode) {
	var bf bytes.Buffer
	first := true
	for node != nil {
		if first {
			first = false
			bf.WriteString(strconv.Itoa(node.Val))
		} else {
			bf.WriteString(fmt.Sprintf(" -> %d", node.Val))
		}
		node = node.Next
	}
	fmt.Println(bf.String())
}
