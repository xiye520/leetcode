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

// 8.字符串转换整数
func TestMyAtoi(t *testing.T) {
	strs := []string{"   -42", "-91283472332", "+", "-", "3.14159", "+1", "+-5", "+5-", "   +0 123", "20000000000000000000", "-   234", "+   234"}
	for _, str := range strs {
		n := MyAtoi(str)
		t.Logf("str: %s转换为数字后为：%d\n", str, n)
	}
}

// 15.三数之和
func TestThreeSum3(t *testing.T) {
	//nums := []int{-1, 0, 1, 2, -1, -4}
	//nums := []int{1, -1, -1, 0}
	nums := []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}
	r := ThreeSum3(nums)
	for _, v := range r {
		t.Log(v)
	}

	t.Log("---------")
	t.Log(r)

}
