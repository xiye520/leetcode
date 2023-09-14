package medium

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
