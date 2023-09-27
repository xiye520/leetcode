package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
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
			args{[]int{1, 1, 2}},
			[][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
		{
			"",
			args{[]int{0, 1}},
			[][]int{{0, 1}, {1, 0}},
		},
		{
			"",
			args{[]int{1}},
			[][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, permuteUnique(tt.args.nums), "permuteUnique(%v)", tt.args.nums)
		})
	}
}
