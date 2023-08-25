package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]int{2, 3, 1, 1, 4}},
			2,
		},
		{
			"",
			args{[]int{2, 3, 0, 1, 4}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, jump(tt.args.nums), "jump(%v)", tt.args.nums)
		})
	}
}
