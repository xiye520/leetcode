package list

//线性表中的链式存储结构
//第一个节点为头节点，并不真实保存数据，头节点基本代表了整个链表

import (
	"fmt"
)

type Elem int

type LinkNode struct {
	Data Elem
	Next *LinkNode
}

//生成头节点
func NewOld() *LinkNode {
	//下面的data可以用来表示链表的长度
	return &LinkNode{0, nil}
}

//在链表的第i个位置前插入一个元素e，复杂度为o(n)
func (head *LinkNode) Insert(i int, e Elem) bool {
	p := head
	j := 1
	for nil != p && j < i {
		p = p.Next
		j++
	}
	if nil == p || j > i {
		fmt.Println("pls check i:", i)
		return false
	}
	s := &LinkNode{Data: e}
	s.Next = p.Next
	p.Next = s
	return true
}

//遍历链表
func (head *LinkNode) Traverse() {
	point := head.Next
	for nil != point {
		fmt.Println(point.Data)
		point = point.Next
	}
	fmt.Println("--------done----------")
}

//删除链表中第i个节点，复杂度为o(n)
func (head *LinkNode) Delete(i int) bool {
	p := head
	j := 1
	for nil != p && j < i {
		p = p.Next
		j++
	}
	if nil == p || j > i {
		fmt.Println("pls check i:", i)
		return false
	}

	p.Next = p.Next.Next
	return true
}

// 获取链表中的第i个元素，复杂度为o(n)
func (head *LinkNode) Get(i int) Elem {
	p := head.Next
	for j := 1; j < i; j++ {
		if nil == p {
			//表示返回错误
			return -100001
		}
		p = p.Next
	}

	return p.Data
}

func main() {
	linkedList := NewOld()
	linkedList.Insert(1, 9)
	linkedList.Insert(1, 99)
	linkedList.Insert(1, 999)
	linkedList.Insert(1, 9999)
	linkedList.Insert(1, 99999)
	linkedList.Insert(1, 999999)
	linkedList.Traverse()

	linkedList.Delete(4)
	linkedList.Traverse()

	e := linkedList.Get(4)
	fmt.Println(e)
}
