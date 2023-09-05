package qsort

func QuickSort(nums []int) []int {
	quicksort1(nums, 0, len(nums)-1)
	return nums
}

func quicksort1(nums []int, left, right int) {
	if left < right {
		i, j := left, right
		pivot := nums[(left+right)/2]

		for {
			if nums[i] < pivot {
				i++
			}
			if nums[j] > pivot {
				j--
			}
			if i >= j {
				break
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		quicksort1(nums, left, i)
		quicksort1(nums, j+1, right)
	}
}

func QuickSort2(nums []int) []int {
	quicksort2(nums, 0, len(nums)-1)
	return nums
}

func quicksort2(nums []int, left, right int) {}
