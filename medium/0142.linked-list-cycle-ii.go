package medium

import "github.com/xiye520/leetcode/util"

// 1.使用map时空复杂度均为O(n)
func detectCycle(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	m := make(map[*util.ListNode]struct{}, 2)
	m[head] = struct{}{}
	tmp := head.Next
	for tmp != nil {
		if _, ok := m[tmp]; ok {
			return tmp
		}
		m[tmp] = struct{}{}
		tmp = tmp.Next
	}

	return nil
}

func detectCycle2(head *util.ListNode) *util.ListNode {
	// https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/linked-list-cycle-ii-kuai-man-zhi-zhen-shuang-zhi-/
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	for true {
		if fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
			if fast == slow {
				break
			}
		} else {
			return nil
		}
	}
	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return fast
}
