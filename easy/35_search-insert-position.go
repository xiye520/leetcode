package easy

/*
35. 搜索插入位置

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

输入: [1,3,5,6], 5
输出: 2
示例 2:

输入: [1,3,5,6], 2
输出: 1
示例 3:

输入: [1,3,5,6], 7
输出: 4
示例 4:

输入: [1,3,5,6], 0
输出: 0
*/

func searchInsert(nums []int, target int) int {
	//fmt.Printf("%s\t%s\t%s\t\n", "mid", "min", "max")
	min, max := 0, len(nums)-1
	for min <= max {
		mid := (min + max) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			max = mid - 1
		} else if nums[mid] < target {
			min = mid + 1
		}
		//fmt.Println(mid, min, max)
	}

	return max + 1
}

func searchInsert2(nums []int, target int) int {
	// 没有把i放入for语句中
	// 是为了兼容，len(nums) == 0 和 target > nums[len(nums)-1]两种情况
	i := 0
	for i < len(nums) && nums[i] <= target {
		// 相等的时候，直接返回
		if nums[i] == target {
			return i
		}
		// 否则，就去检查下一个
		i++
	}
	return i
}
