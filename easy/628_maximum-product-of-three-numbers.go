package easy

import "math"

// 只要求出数组中最大的三个数以及最小的两个数，因此我们可以不用排序，用线性扫描直接得出这五个数。
func maximumProduct(nums []int) int {
	min1, min2 := math.MaxInt32, math.MaxInt32
	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
	for _, n := range nums {
		if n <= min1 {
			min2 = min1
			min1 = n
		} else if n <= min2 { // n lies between min1 and min2
			min2 = n
		}
		if n >= max1 { // n is greater than max1, max2 and max3
			max3 = max2
			max2 = max1
			max1 = n
		} else if n >= max2 { // n lies betweeen max1 and max2
			max3 = max2
			max2 = n
		} else if n >= max3 { // n lies betwen max2 and max3
			max3 = n
		}
	}

	return max(max1*max2*max3, max1*min1*min2)
}
