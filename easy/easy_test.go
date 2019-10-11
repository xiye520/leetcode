package easy

import (
	"fmt"
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

// 35.搜索插入位置
func Test_searchInsert(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	t.Log(searchInsert(nums, target))
	nums = []int{1, 3, 5, 6}
	t.Log(searchInsert(nums, target))

}

func Test_searchInsert2(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	t.Log(searchInsert2(nums, target))
	nums = []int{1, 3, 5, 6}
	t.Log(searchInsert2(nums, target))
}

func Example_searchInsert() {
	fmt.Println(searchInsert([]int{-1, 0, 3, 5, 9, 12}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert2([]int{-1, 0, 3, 5, 9, 12}, 2))
	fmt.Println(searchInsert2([]int{1, 3, 5, 6}, 2))

	// Output:
	// 2
	// 1
	// 2
	// 1
}

// 70.爬楼梯
func Test_climbStairs(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(i, climbStairs(i))
	}

	for i := 0; i < 10; i++ {
		t.Log(i, climbStairs2(i))
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

// 704.二分查找
func Test_binarySearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 3
	t.Log(search(nums, target))
}
