package qsort

import (
	"reflect"
	"testing"
)

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
			"case 1",
			args{[]int{12, 84, 5, 63, 37, 9, 56, 1399}},
			[]int{5, 9, 12, 37, 56, 63, 84, 1399},
		},
		{
			"case 2",
			args{[]int{1, 2, 3, 4, 5}},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"case 3",
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
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case 1",
			args{[]int{12, 84, 5, 63, 37, 9, 56, 1399}},
			[]int{5, 9, 12, 37, 56, 63, 84, 1399},
		},
		{
			"case 2",
			args{[]int{1, 2, 3, 4, 5}},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"case 3",
			args{[]int{5, 4, 3, 2, 1}},
			[]int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort2() = %v, want %v", got, tt.want)
			}
		})
	}
}
