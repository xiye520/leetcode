package easy

import (
	"testing"
)

func Test_findLengthOfLCIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1, 3, 5, 4, 7",
			args{[]int{1, 3, 5, 4, 7}},
			3,
		},
		{
			"1, 3, 5, 6, 7",
			args{[]int{1, 3, 5, 6, 7}},
			5,
		},
		{
			"",
			args{[]int{7, 3, 2, 1, 0}},
			1,
		},
		{
			"1, 3, 5, 6, 7",
			args{[]int{1, 3, 5, 6, 4}},
			4,
		},
		{
			"1,3,5,4,2,3,4,5",
			args{[]int{1, 3, 5, 4, 2, 3, 4, 5}},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfLCIS(tt.args.nums); got != tt.want {
				t.Errorf("findLengthOfLCIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
