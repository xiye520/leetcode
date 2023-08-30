package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"",
			args{[]int{-1, 0, 1, 2, -1, -4}},
			[][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ThreeSum2(tt.args.nums), "ThreeSum2(%v)", tt.args.nums)
		})
	}
}

// 15.三数之和
func TestThreeSum3(t *testing.T) {
	//nums := []int{-1, 0, 1, 2, -1, -4}
	//nums := []int{1, -1, -1, 0}
	nums := []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}
	r := ThreeSum3(nums)
	for _, v := range r {
		t.Log(v)
	}

	t.Log("---------")
	t.Log(r)

}

func Test_sortSlice(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"",
			args{[]int{-1, 0, 1, 2, -1, -4}},
			[]int{-4, -1, -1, 0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, sortSlice(tt.args.nums), "sortSlice(%v)", tt.args.nums)
		})
	}
}
