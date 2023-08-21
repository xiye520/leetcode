package main

import (
	"reflect"
	"testing"
)

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"20200202",
			args{"20200202"},
			true,
		},
		{
			"21211212",
			args{"21211212"},
			false,
		},
		{
			"19700101",
			args{"19700101"},
			false,
		},
		{
			"22222222",
			args{"22222222"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transfer(t *testing.T) {
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
			"1,2,3,4,5,6,7",
			args{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			"1,2,3",
			args{[]int{1, 2, 3}, 4},
			[]int{3, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transfer(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transfer() = %v, want %v", got, tt.want)
			}
		})
	}
}
