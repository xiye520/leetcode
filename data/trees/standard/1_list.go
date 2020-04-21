package main

import (
	"container/list"
	"fmt"
)

/*
操作List
(1)直接使用container/list包下的New()新建一个空的List

添加，遍历，取首尾，取中间元素
*/
func main() {
	//实例化
	mylist := list.New()
	fmt.Println(mylist)

	//添加
	mylist.PushFront("a") //["a"]
	mylist.PushBack("b")  //["a","b"]
	mylist.PushBack("c")  //["a","b","c"]
	//在最后一个元素的前面添加
	mylist.InsertBefore("d", mylist.Back()) //["a","b","d","c"]
	mylist.InsertAfter("e", mylist.Front()) //["a","e","b","d","c"]

	//遍历
	for e := mylist.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ") //a e b d c
	}
	fmt.Println("")

	//取首尾
	fmt.Println(mylist.Front().Value) //a
	fmt.Println(mylist.Back().Value)  //c

	//取中间的元素，通过不断的Next()
	n := 3
	var curr *list.Element
	if n > 0 && n <= mylist.Len() {
		if n == 1 {
			curr = mylist.Front()
		} else if n == mylist.Len() {
			curr = mylist.Back()
		} else {
			curr = mylist.Front()
			for i := 1; i < n; i++ {
				curr = curr.Next()
			}
		}
	} else {
		fmt.Println("n的数值不对")
	}
	fmt.Println(curr.Value) //b
}
