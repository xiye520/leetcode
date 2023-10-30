package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumber260(t *testing.T) {
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
			args{[]int{-1, 0}},
			[]int{-1, 0},
		},
		{
			"",
			args{[]int{1, 2, 1, 3, 2, 5}},
			[]int{3, 5},
		},
		{
			"",
			args{[]int{1, 0}},
			[]int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, singleNumber260(tt.args.nums), "singleNumber260(%v)", tt.args.nums)
		})
	}
}
