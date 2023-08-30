package medium

import (
	"github.com/xiye520/leetcode/util"
	"testing"
)

// 8.字符串转换整数
func TestMyAtoi(t *testing.T) {
	strs := []string{"   -42", "4193 with words", "words and 98", "-91283472332", "+", "-", "3.14159", "+1", "+-5", "+5-", "   +0 123",
		"20000000000000000000", "-   234", "+   234"}
	for _, str := range strs {
		n := myAtoi(str)
		t.Logf("str: %s转换为数字后为：%d\n", str, n)
	}
}

// 8.字符串转换整数2
func Test_myAtoi2(t *testing.T) {
	strs := []string{"   -42", "4193 with words", "words and 98", "-91283472332", "+", "-", "3.14159", "+1", "+-5", "+5-", "   +0 123",
		"20000000000000000000", "-   234", "+   234"}
	for _, str := range strs {
		n := myAtoi2(str)
		t.Logf("str: %s转换为数字后为：%d\n", str, n)
	}
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
func Test_searchRange(t *testing.T) {
	arr := []int{5, 7, 7, 8, 8, 10}
	res := searchRange(arr, 10)
	t.Log(res)
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
func Test_searchRange2(t *testing.T) {
	arr := []int{5, 7, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 10}
	{
		target := 10
		res := searchRange2(arr, target)
		t.Log(target, res)
	}

	{
		target := 8
		res := searchRange2(arr, target)
		t.Log(target, res)
	}

	{
		target := 7
		res := searchRange2(arr, target)
		t.Log(target, res)
	}

	{
		target := 6
		res := searchRange2(arr, target)
		t.Log(target, res)
	}
}

// 78.子集
func Test_subsets(t *testing.T) {
	nums := []int{1, 2, 3}
	nums = []int{1, 2, 3, 4}
	res := subsets(nums)
	for _, r := range res {
		t.Log(r)
	}
}

// 92. 反转链表 II
func Test_reverseBetween(t *testing.T) {
	util.Sord3()
}

func Test_medium(t *testing.T) {
	nums := []int{1, 2, 3}
	nums = []int{1, 2, 3, 4}
	res := subsets(nums)
	for _, r := range res {
		t.Log(r)
	}
}
