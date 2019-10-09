package easy

import (
	"github.com/stretchr/testify/assert"
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

// 9.回文数字
func Test_isPalindrome(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: 12321,
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: 1231,
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: -12321,
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isPalindrome(p.one), "输入:%v", p)
	}
}

// 70.爬楼梯
func Test_climbStairs1(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(i, climbStairs(i))
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
