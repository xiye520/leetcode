package easy

import (
	"testing"
)

// 1.两数之和
func TestTwoSum(t *testing.T) {
	target := 9
	nums := []int{2, 7, 11, 15}

	indexs := twoSum(nums, target)

	if len(nums) < 2 || nums[indexs[0]]+nums[indexs[1]] != target {
		t.Error("fails....")
	}
}
