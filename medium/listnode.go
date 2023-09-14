package medium

import (
	"bytes"
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
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
