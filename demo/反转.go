package main

import "log"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			log.Printf("不用删除的索引：%d，当前标记：%d", i, j)
			if i != j {
				log.Println("待交换：", i, j)
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
			log.Println("索引：", j)
		}
	}
	log.Println(nums, j)
	return j
}
