package _828

import (
	"fmt"
)

func ExampleNewStack() {
	s := NewStack()
	s.push(1)
	s.push(2)
	s.push(3)
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())

	// output:
	// 1
	// 2
	// 3
}
