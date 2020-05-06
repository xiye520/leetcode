package easy

// 实现一个栈

type MyStack struct {
	nums []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{nums: []int{}}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.nums = append(this.nums, x)
	return
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	ret := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	return ret
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.nums[len(this.nums)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.nums) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
