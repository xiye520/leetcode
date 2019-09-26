package medium

import (
	"fmt"
)

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) prinfListInReserveOrder() {
	if l.Next != nil {
		l.Next.prinfListInReserveOrder()
	}
	fmt.Print("<-  ", l.Val)
}

func (l ListNode) printListInOrder() {
	fmt.Print("  ->", l.Val)
	if l.Next != nil {
		l.Next.printListInOrder()
	}
}

func (l *ListNode) Push(v int) {
	TempList := l
	for TempList.Next != nil {
		TempList = TempList.Next
	}
	TempList.Next = &ListNode{
		Val: v,
	}
}

func (l *ListNode) Set(v int) {
	l.Val = v
}

func sord3() {
	l := &ListNode{
		Val: 123,
	}
	for i := 0; i < 10; i++ {
		l.Push(i)
	}
	fmt.Println("顺序打印：")
	l.printListInOrder()
	fmt.Println("")
	fmt.Println("倒序打印：")
	l.prinfListInReserveOrder()
}
