package medium

import (
	"math"
	"strconv"
)

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	length := func(l *ListNode) (n int) {
		n = 1
		for l.Next != nil {
			n++
			l = l.Next
		}
		return n
	}

	last := func(l *ListNode) (n *ListNode) {
		if l.Next == nil {
			return l
		}

		for l.Next != nil {
			n = l.Next
		}

		return n
	}
	n1 := length(l1)
	n2 := length(l2)

	var a1 int
	for i := 0; i < n1; i++ {
		a1 += l1.Val * int(math.Pow10(n1-i-1))
	}
	var b1 int
	for i := 0; i < n2; i++ {
		b1 += l2.Val * int(math.Pow10(n2-i-1))
		l2 = l2.Next
	}

	sum := a1 + b1 //807
	r := &ListNode{
		Val:  sum % 10,
		Next: nil,
	}
	for sum/10 > 0 {
		sum = sum / 10
		n := last(r)
		n.Next = &ListNode{
			Val: sum % 10,
		}
	}

	return r
}

//func (l *ListNode) Length() (n int) {
//	n = 1
//	for l.Next != nil {
//		n++
//		l = l.Next
//	}
//	return n
//}
//
//func (l *ListNode) Last() (n *ListNode) {
//	for l.Next != nil {
//		n = l.Next
//	}
//
//	return n
//}

func setData(l *ListNode, len int, str string) (*ListNode, int, string) {
	if len == 0 {
		return l, len, str
	}

	v, _ := strconv.Atoi(str[:1])
	next := &ListNode{
		Val:  v,
		Next: nil,
	}

	return setData(next.Next, len-1, str[1:len])
}

//func A(l1 *ListNode, l2 *ListNode) *ListNode {
//	a := 100*l1.Val + 10*l1.Next.Val + l1.Next.Next.Val
//	b := 100*l2.Val + 10*l2.Next.Val + l2.Next.Next.Val
//	sum := a + b //807
//
//	return &ListNode{
//		Val: sum % 10,
//		Next: &ListNode{
//			Val: (sum % 100) / 10,
//			Next: &ListNode{
//				Val: sum / 100,
//				//Next:nil
//			},
//		},
//	}
//}
