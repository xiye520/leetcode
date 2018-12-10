package easy

import (
	"log"
	"testing"
)

/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

func ThreeSum2(nums []int) [][]int {
	nums = sort(nums)
	var r [][]int
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					var new bool
					for h := 0; h < len(r); h++ {
						if (nums[i] == r[h][0] && nums[j] == r[h][1]) || (nums[i] == r[h][0] && nums[j] == r[h][2]) || (nums[i] == r[h][1] && nums[j] == r[h][2]) || (nums[i] == r[h][1] && nums[j] == r[h][0]) || (nums[i] == r[h][2] && nums[j] == r[h][0]) || (nums[i] == r[h][2] && nums[j] == r[h][1]) {
							new = true
							break
						}
					}
					if !new {
						r = append(r, []int{nums[i], nums[j], nums[k]})
					}
				}
			}
		}
	}

	return r
}

func sort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	log.Println(nums)

	return nums
}

//func TestThreeSum2(t *testing.T) {
//	//nums := []int{-1, 0, 1, 2, -1, -4}
//	nums := []int{1, -1, -1, 0}
//	r := ThreeSum2(nums)
//	for _, v := range r {
//		t.Log(v)
//	}
//
//	t.Log("---------")
//	t.Log(r)
//
//}

func ThreeSum3(nums []int) [][]int {
	nums = sort(nums)
	var r [][]int
	l := len(nums)
	for i := 0; i < l-2; i++ {
		j, k := i+1, l-1
		if i > 0 && nums[i] == nums[i-1] { // 去重
			continue
		}

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
				for j < k && nums[j] == nums[j-1] { //去重
					j++
				}
			} else if sum > 0 {
				k--
				for j < k && nums[k] == nums[k+1] { //去重
					k--
				}
			} else {
				r = append(r, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] { //去重
					j++
				}
				for j < k && nums[k] == nums[k+1] { //去重
					k--
				}
			}
		}
	}

	return r
}

func TestThreeSum3(t *testing.T) {
	//nums := []int{-1, 0, 1, 2, -1, -4}
	nums := []int{1, -1, -1, 0}
	r := ThreeSum3(nums)
	for _, v := range r {
		t.Log(v)
	}

	t.Log("---------")
	t.Log(r)

}
