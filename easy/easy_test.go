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

	indexs := twoSum2(nums, target)

	if len(nums) < 2 || nums[indexs[0]]+nums[indexs[1]] != target {
		t.Fatal("fails....")
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

// 66.加一
func Test_plusOne(t *testing.T) {
	nums := []int{9, 9, 9, 9}

	t.Log(plusOne(nums))
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

// 121. 买卖股票的最佳时机
func Test_maxProfit(t *testing.T) {
	prices := []int{1, 2, 6, 4, 5}
	prices = []int{7, 1, 5, 3, 2, 6}
	t.Log(maxProfit(prices))
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

// 326.3的幂
func Example_isPowerOfThree() {
	fmt.Println(isPowerOfThree(0))
	fmt.Println(isPowerOfThree(1))  // true
	fmt.Println(isPowerOfThree(3))  // true
	fmt.Println(isPowerOfThree(9))  // true
	fmt.Println(isPowerOfThree(27)) // true
	fmt.Println(isPowerOfThree(28))
	fmt.Println(isPowerOfThree(29))
	fmt.Println(isPowerOfThree(30))
	fmt.Println(isPowerOfThree(45))

	// Output:
	// false
	// true
	// true
	// true
	// true
	// false
	// false
	// false
	// false
}

// 704.二分查找
func Test_binarySearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 3
	t.Log(search(nums, target))
}

/*
示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。
*/

func Test_easy(t *testing.T) {
	nums := []int{9, 9, 9, 9}

	t.Log(plusOne(nums))
}

// go test -v -test.run Test_kidsWithCandies
func Test_kidsWithCandies(t *testing.T) {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	res := kidsWithCandies(candies, extraCandies)
	want := []bool{true, true, true, false, true}
	for i := range res {
		if res[i] != want[i] {
			t.Fatal(i, res[i])
		}
	}
	t.Log("success")
}
