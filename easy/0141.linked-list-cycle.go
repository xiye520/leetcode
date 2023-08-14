package easy

import (
	"github.com/xiye520/leetcode/util"
)

/** 141. 环形链表
给你一个链表的头节点 head ，判断链表中是否有环。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。
如果链表中存在环 ，则返回 true 。 否则，返回 false 。

输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。

输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。

输入：head = [1], pos = -1
输出：false
解释：链表中没有环。
*/

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
