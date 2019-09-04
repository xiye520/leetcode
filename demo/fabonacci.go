package main

import "fmt"

// 1.闭包实现
func fabonacci_pack() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

// 2.递归实现
func fabo_recursion(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fabo_recursion(n-1) + fabo_recursion(n-2)
}

// 3.channel实现
func fabo_channel(output chan<- int, quit chan bool) {
	x, y := 0, 1
	for {
		select {
		case output <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("stop fabonacci ")
			return
		}
	}

}

func main() {
	//f := fabonacci_pack()
	//for i := 0; i < 10; i++ {
	//10fmt.Println(i, f())
	//}

	//for i := 0; i < 10; i++ {
	//10fmt.Println(i, fabo_recursion(i))
	//}

	quit := make(chan bool)
	output := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			n := <-output
			fmt.Println(n)
		}
		quit <- true
	}()

	fabo_channel(output, quit)
	//close(quit)
}
