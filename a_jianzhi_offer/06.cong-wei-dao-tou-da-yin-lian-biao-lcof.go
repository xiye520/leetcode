package jianzhi_offer

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	return appendData(head)
}

func appendData(head *ListNode) []int {
	if head.Next != nil {
		list := appendData(head.Next)
		list = append(list, head.Val)
		return list
	}

	return []int{head.Val}
}

func reversePrint2(head *ListNode) []int {
	head = reverseList(head)
	res := make([]int, 0, 100)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	var prev *ListNode
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}

	return prev
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		// 将下一个节点指向头节点
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}
