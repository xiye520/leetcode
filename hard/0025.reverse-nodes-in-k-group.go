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
