package sort

import (
	"fmt"
	"reflect"
	"testing"
)

/*
1，测试单个文件，一定要带上被测试的原文件
    go test -v  wechat_test.go wechat.go

2，测试单个方法
    go test -v -test.run TestRefreshAccessToken

go test -test.bench=".*"
*/

var (
	array []int = []int{0, 20, 10, 25, 15, 30, 28, 55, 432, 432432, 4234, 333, 333, 21, 22, 3, 30, 8, 20, 2, 7, 9, 50, 80, 1, 4}
)

// 冒泡1
func TestBubbleSort(t *testing.T) {
	//array := []int{12, 84, 5, 63, 37, 9, 56, 1399}
	fmt.Println("befor sort array:", array)
	BubbleSort(array)
	fmt.Println("after sort array:", array)
}

// 冒泡2
func Test_bubbleSort2(t *testing.T) {
	fmt.Println("befor sort array:", array)
	bubbleSort2(array)
	fmt.Println("after sort array:", array)
}

// 冒泡3
func Test_bubbleSort3(t *testing.T) {
	fmt.Println("befor sort array:", array)
	bubbleSort3(array)
	fmt.Println("after sort array:", array)
}

/*
性能分析对比：
goos: windows
goarch: 386
pkg: leetcode/data_structures/sort
BenchmarkBubbleSort-2            2000000               904 ns/op
BenchmarkBubbleSort2-2          20000000                71.2 ns/op
BenchmarkBubbleSort3-2          20000000                96.4 ns/op
*/
func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(array)
	}
}

func BenchmarkBubbleSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubbleSort2(array)
	}
}

func BenchmarkBubbleSort3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubbleSort3(array)
	}
}

func TestSelectSort(t *testing.T) {
	//array := []int{12, 84, 5, 63, 37, 9, 56, 1399}
	fmt.Println("befor sort array:", array)
	SelectSort(array)
	fmt.Println("after sort array:", array)
}

func TestQuickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"QuickSort",
			args{[]int{5, 4, 3, 2, 1}},
			[]int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSort2(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"QuickSort2",
			args{[]int{5, 4, 3, 2, 1}},
			[]int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort2(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
