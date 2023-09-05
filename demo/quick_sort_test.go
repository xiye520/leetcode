package main

import (
	"reflect"
	"testing"
)

func TestQuickSort6(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"success",
			args{[]int{5, 3, 2, 4, 1}},
			[]int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort6(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort6() = %v, want %v", got, tt.want)
			}
		})
	}
}
