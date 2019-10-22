package medium

import "fmt"

/*
55. 跳跃游戏

给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个位置。

示例 1:
输入: [2,3,1,1,4]
输出: true
解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。

示例 2:
输入: [3,2,1,0,4]
输出: false
解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
*/
func canJump(nums []int) bool {
	/*
		如果所有元素都不为0， 那么一定可以跳到最后；
		从后往前遍历，如果遇到nums[i] = 0，就找i前面的元素j，使得nums[j] > i - j。如果找不到，则不可能跳跃到num[i+1]，返回false。
	*/
	if len(nums) > 1 && nums[0] == 0 {
		return false
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] == 0 {
			can := false
			for j := i; j >= 0; j-- {
				//fmt.Println(nums[j], i, j)
				if nums[j] > i-j {
					i = j
					can = true
					break
				}
			}
			if !can {
				return false
			}
		}
	}

	return true
}

func canJump2(nums []int) bool {
	/*
		解题思路：
		1.如果某一个作为 起跳点 的格子可以跳跃的距离是 3，那么表示后面 3 个格子都可以作为 起跳点。
		2.可以对每一个能作为 起跳点 的格子都尝试跳一次，把 能跳到最远的距离 不断更新。
		3.如果可以一直跳到最后，就成功了。
	*/
	maxIndex := 0
	for i := 0; i < len(nums); i++ {
		if i > maxIndex {
			return false
		}
		maxIndex = max(maxIndex, i+nums[i])
	}

	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
[2,3,1,1,4] 		true
[3,2,1,0,4] 		false
[0,3,1,0,4,3,2,0,1] false
[0] 				true
[2,0,0,3,0,1,4,4] 	false
*/
func test_canJump() {
	//nums := []int{2, 0, 0, 3, 0, 1, 4, 4}
	fmt.Println(canJump2([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump2([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump2([]int{0, 3, 1, 0, 4, 3, 2, 0, 1}))
	fmt.Println(canJump2([]int{0}))
	fmt.Println(canJump2([]int{2, 0, 0, 3, 0, 1, 4, 4}))
}
