package medium

import "testing"

func TestGetNodeLen(t *testing.T) {
	sum := 235
	a := &ListNode{
		Val: sum % 10,
		Next: &ListNode{
			Val: (sum % 100) / 10,
			//Next: &ListNode{
			//	Val: sum / 100,
			//	//Next:nil
			//},
		},
	}
	l := getNodeLen(a)
	t.Log(l)
}
