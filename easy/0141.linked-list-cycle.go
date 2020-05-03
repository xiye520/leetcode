package easy

import "leetcode/util"

// 使用map，时间复杂度、空间复杂度均为O(n)
func hasCycle2(head *util.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	m := make(map[*util.ListNode]bool, 5)
	for head.Next != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = true
		head = head.Next
	}
	return false
}

// 快慢指针，最优解，时间复杂度O(n),空间复杂度O(1)
func hasCycle(head *util.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
