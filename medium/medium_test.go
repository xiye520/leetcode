package medium

import "testing"

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

// 11.盛最多水的容器
func Test_maxArea(t *testing.T) {
	result := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	if result != 49 {
		t.Fatalf("maxArea not match 49,it is: %d", result)
	}
	t.Logf("maxArea: %d", result)
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
