package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"2",
			args{[]int{1, 2, 3, 4, 5}, 2},
			[]int{2, 1, 4, 3, 5},
		},
		{
			"3",
			args{[]int{1, 2, 3, 4, 5}, 3},
			[]int{3, 2, 1, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKNumbers(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleTest1() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	reverseKNumbers(arr, 2)
	fmt.Println(arr)
	head := buildListNode(arr)
	printList(head)

	// output:
	// [1 2 3 4 5]
	// [2 1 4 3 5]
	// 2 -> 1 -> 4 -> 3 -> 5
}
