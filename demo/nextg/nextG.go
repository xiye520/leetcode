package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(" X: ", i)
			wg.Done()
		}()
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(" Y: ", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

/**
nextG，新的goroutine顺序放入p的本地runnable queue
Y:  9
X:  10
X:  10
X:  10
X:  10
X:  10
X:  10
X:  10
X:  10
X:  10
X:  10
Y:  0
Y:  1
Y:  2
Y:  3
Y:  4
Y:  5
Y:  6
Y:  7
Y:  8

*/
