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

func QuickSort2(arr []int) []int {
	quicksort2(arr, 0, len(arr)-1)
	return arr
}

func quicksort2(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partiton(arr, left, right)
	quicksort2(arr, left, pivot-1)
	quicksort2(arr, pivot+1, right)
}

func partiton(arr []int, left, right int) int {
	key := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] < key {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}

	}

	arr[i], arr[right] = arr[right], arr[i]
	return i
}
