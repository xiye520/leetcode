package easy

import "testing"

/*
测试整个文件：
go test -v 70_climbing-stairs_test.go -test.bench=".*"

测试单个函数：
go test -v 70_climbing-stairs_test.go -test.run Benchmark_climbStairs

go test -test.bench=".*"
*/
// 70.爬楼梯
func Test_climbStairs2(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(i, climbStairs(i))
	}

	for i := 0; i < 10; i++ {
		t.Log(i, climbStairs2(i))
	}
}

func Benchmark_climbStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		climbStairs(i)
	}
}

func Benchmark_climbStairs2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		climbStairs2(i)
	}
}
