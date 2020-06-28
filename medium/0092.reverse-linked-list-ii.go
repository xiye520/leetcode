package medium

/*
92. 反转链表 II

反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m >= n {
		return head
	}

	var dump ListNode // 定义结果链表
	dump.Next = head  // 初始化
	start := &dump    // 更新结果链表
	for i := 1; i < m; i++ {
		start = start.Next
	}

	end := start.Next
	for i := m; i < n; i++ {
		cur := end.Next          // 更新当前节点
		end.Next = end.Next.Next // 更新下一节点

		cur.Next = start.Next // 更新当前节点的下一节点
		start.Next = cur
	}

	return dump.Next
}
