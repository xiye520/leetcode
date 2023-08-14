package medium

import "math/rand"

/** 380. O(1) 时间插入、删除和获取随机元素
实现RandomizedSet 类：

RandomizedSet() 初始化 RandomizedSet 对象
bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。

示例：
输入
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
输出
[null, true, false, true, 2, true, false, 2]

解释
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
randomizedSet.remove(2); // 返回 false ，表示集合中不存在 2 。
randomizedSet.insert(2); // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
randomizedSet.getRandom(); // getRandom 应随机返回 1 或 2 。
randomizedSet.remove(1); // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
randomizedSet.insert(2); // 2 已在集合中，所以返回 false 。
randomizedSet.getRandom(); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。

https://leetcode.cn/problems/insert-delete-getrandom-o1/?envType=study-plan-v2&envId=top-interview-150
*/

type RandomizedSet struct {
	nums []int
	m    map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, make(map[int]int)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	// 存元素对应的在数组中索引
	this.m[val] = len(this.nums)
	// 存元素
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	id, ok := this.m[val]
	if !ok {
		return false
	}
	// 要剔除的数跟最后一个元素互换位置
	this.m[this.nums[len(this.nums)-1]] = id
	this.nums[id] = this.nums[len(this.nums)-1]
	delete(this.m, val)
	// 剔除最后一个元素
	this.nums = this.nums[:len(this.nums)-1]
	return true

}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
