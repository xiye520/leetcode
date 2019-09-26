package main

import (
	"fmt"
	"strconv"
)

type LinkedList struct {
	Value string
	Next  *LinkedList
}

func (ll LinkedList) prinfListInReserveOrder() {
	if ll.Next != nil {
		ll.Next.prinfListInReserveOrder()
	}
	fmt.Print("<-  ", ll.Value)
}

func (ll LinkedList) prinfListInOrder() {
	fmt.Print("  ->", ll.Value)
	if ll.Next != nil {
		ll.Next.prinfListInOrder()
	}
}

func (ll *LinkedList) Push(v string) {
	TempList := ll
	for TempList.Next != nil {
		TempList = TempList.Next
	}
	TempList.Next = &LinkedList{
		Value: v,
	}
}

func (ll *LinkedList) Set(v string) {
	ll.Value = v
}

func main() {
	ll := &LinkedList{
		Value: "123",
	}
	for i := 0; i < 10; i++ {
		ll.Push(strconv.Itoa(i))
	}
	fmt.Println("PinrfInOrder")
	ll.prinfListInOrder()
	fmt.Println("")
	fmt.Println("PinrfInReservedOrder")
	ll.prinfListInReserveOrder()
}
