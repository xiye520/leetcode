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

// 169.求众数
func Test_majorityElement(t *testing.T) {
	arr := []int{2, 2, 1, 1, 1, 2, 2}
	result := majorityElement(arr)
	t.Log(result)
}

// 169.求众数2
func Test_majorityElement2(t *testing.T) {
	arr := []int{2, 2, 1, 1, 1, 2, 2}
	result := majorityElement2(arr)
	t.Log(result)
}

// 169.求众数3
func Test_majorityElement3(t *testing.T) {
	arr := []int{2, 2, 1, 1, 1, 2, 2}
	result := majorityElement3(arr)
	t.Log(result)
}
