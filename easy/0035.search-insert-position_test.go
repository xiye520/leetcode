package easy

import (
	"fmt"
	"testing"
)

// 35.搜索插入位置
func Test_searchInsert(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	t.Log(searchInsert(nums, target))
	nums = []int{1, 3, 5, 6}
	t.Log(searchInsert(nums, target))

}

// 35.搜索插入位置
func Test_searchInsert2(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	t.Log(searchInsert2(nums, target))
	nums = []int{1, 3, 5, 6}
	t.Log(searchInsert2(nums, target))
}

// 35.搜索插入位置
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
