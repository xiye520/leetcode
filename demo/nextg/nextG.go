package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
** 考点：maxProcs会不会上下文切换、nextG、loop var

1、当wg的数量是20，因为基本上1ms就跑完了，来不及触发 10ms 的运行时强制调度

2、当wg的数量增加到2000，第一个输出的永远是Y: count/2-1，剩下的就是随机输出了
*/
func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	count := 20
	wg.Add(count)
	for i := 0; i < count/2; i++ {
		go func() {
			fmt.Print(" X: ", i)
			wg.Done()
		}()
	}

	for i := 0; i < count/2; i++ {
		go func(i int) {
			fmt.Print(" Y: ", i)
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
