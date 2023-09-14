package _828

import "sync"

// 题目1️：用两个栈实现队列
//题目描述：请你用两个栈实现一个队列，完成队列的Push和Pop操作。 队列中的元素为int类型。
//示例:
//输入:
//push(1)
//push(2)
//push(3)
//输出:
//pop() // 返回1
//pop() // 返回2
//push(4) pop() // 返回3
//pop() // 返回4

type Stack struct {
	in  []int
	out []int
	rw  sync.RWMutex
}

// NewStack 初始化
func NewStack() Stack {
	return Stack{make([]int, 0, 1024), make([]int, 0, 1024), sync.RWMutex{}}
}

// 输入
func (s *Stack) push(i int) {
	// 直接压到输入的切片末尾
	s.in = append(s.in, i)
}

// 出栈(如果没数据返回-1)
func (s *Stack) pop() int {
	var res int
	// 先出
	if len(s.out) != 0 {
		res = s.out[len(s.out)-1]
		s.out = s.out[:len(s.out)-1]
		return res
	}
	if len(s.in) != 0 {
		res = s.in[0]
		for i := len(s.in) - 1; i >= 1; i-- {
			s.out = append(s.out, s.in[i])
		}
		s.in = make([]int, 0, 1024)
		return res
	}

	return -1
}
