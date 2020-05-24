package main

import (
	"testing"
)

// go test -v -test.run Test_fabp_dp
func Test_fabp_dp(t *testing.T) {
	//f := fabonacci_pack()
	//for i := 0; i < 10; i++ {
	//10fmt.Println(i, f())
	//}

	//for i := 0; i < 10; i++ {
	//10fmt.Println(i, fabo_recursion(i))
	//}

	//quit := make(chan bool)
	//output := make(chan int)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		n := <-output
	//		fmt.Println(n)
	//	}
	//	quit <- true
	//}()
	//
	//fabo_channel(output, quit)
	//close(quit)
	for i := 0; i < 10; i++ {
		t.Logf("i: %d, fabonacci: %d", i, fabp_dp(i))
	}
}
