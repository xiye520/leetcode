package easy

/*
将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	pre := new(ListNode)
	cur := pre
	for {
		if l1 == nil {
			cur.Next = l2
			break
		}
		if l2 == nil {
			cur.Next = l1
			break
		}

		if l1.Val < l2.Val {
			cur.Next = &ListNode{
				Val: l1.Val,
			}
			l1 = l1.Next
		} else {
			cur.Next = &ListNode{
				Val: l2.Val,
			}
			l2 = l2.Next
		}
		cur = cur.Next
	}

	return pre.Next
}

// 最精简的写法
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if l1 == nil {
		curr.Next = l2
	}
	if l2 == nil {
		curr.Next = l1
	}

	return head.Next
}

func mergeTwoLists3(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}
	if list1 != nil {
		curr.Next = list1
	}
	if list2 != nil {
		curr.Next = list2
	}

	return head.Next
}
