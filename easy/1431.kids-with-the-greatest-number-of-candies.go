package easy

/*
1431. 拥有最多糖果的孩子
给你一个数组 candies 和一个整数 extraCandies ，其中 candies[i] 代表第 i 个孩子拥有的糖果数目。

对每一个孩子，检查是否存在一种方案，将额外的 extraCandies 个糖果分配给孩子们之后，此孩子有 最多 的糖果。注意，允许有多个孩子同时拥有 最多 的糖果数目。



示例 1：

输入：candies = [2,3,5,1,3], extraCandies = 3
输出：[true,true,true,false,true]
*/
func kidsWithCandies(candies []int, extraCandies int) []bool {
	//m := make(map[int]bool, len(candies))
	arr := make([]int, len(candies))
	for i, v := range candies {
		arr[i] = v + extraCandies
	}

	res := make([]bool, len(candies))
	for k, v := range arr {
		//if b, ok := m[v]; ok {
		//	res[k] = b
		//	continue
		//}
		res[k] = true

		for i := 0; i < len(candies); i++ {
			if i == k {
				continue
			}
			if candies[i] > v {
				res[k] = false
				//m[v] = false
				break
			}
		}
		//m[v] = true
	}

	return res
}
