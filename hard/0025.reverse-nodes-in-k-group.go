package hard

func reverseKGroup(head *ListNode, k int) *ListNode {
	arr := make([]int, 0, 10)
	for head != nil {
		arr = append(arr, head.val)
		head = head.Next
	}
	reverseKNumbers(arr, k)
	newHead := buildListNode(arr)

	return newHead
}

func reverseKNumbers(nums []int, k int) []int {
	if len(nums) < k {
		return nums
	}
	group := len(nums) / k
	for i := 0; i < group; i++ {
		start := i * k
		end := i*k + k - 1
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}

	return nums
}

func buildListNode(nums []int) *ListNode {
	var dummy ListNode
	curr := &dummy
	for _, v := range nums {
		curr.Next = &ListNode{val: v}
		curr = curr.Next
	}

	return dummy.Next
}

// https://leetcode.cn/problems/reverse-nodes-in-k-group/solutions/1239394/jian-dan-yi-dong-go-by-dfzhou6-4cha/
func reverseKGroup2(head *ListNode, k int) *ListNode {
	cur := head
	for i := 0; i < k; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}
	newHead := reverse(head, cur)
	head.Next = reverseKGroup(cur, k)
	return newHead
}

func reverse(start, end *ListNode) *ListNode {
	var pre *ListNode
	cur := start
	for cur != end {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

// https://leetcode.cn/problems/reverse-nodes-in-k-group/solutions/248591/k-ge-yi-zu-fan-zhuan-lian-biao-by-leetcode-solutio/
func reverseKGroup3(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}
