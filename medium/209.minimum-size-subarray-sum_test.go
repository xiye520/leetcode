package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"7",
			args{7, []int{2, 3, 1, 2, 4, 3}},
			2,
		},
		{
			"4",
			args{4, []int{1, 4, 4}},
			1,
		},
		{
			"false",
			args{11, []int{1, 1, 1, 1, 1, 1, 1, 1}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minSubArrayLen(tt.args.target, tt.args.nums), "minSubArrayLen(%v, %v)", tt.args.target, tt.args.nums)
		})
	}
}
