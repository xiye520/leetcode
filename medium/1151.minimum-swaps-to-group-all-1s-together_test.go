package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minSwaps(t *testing.T) {
	type args struct {
		in0 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]int{1, 0, 1, 0, 1}},
			1,
		},
		{
			"",
			args{[]int{0, 0, 0, 1, 0}},
			0,
		},
		{
			"",
			args{[]int{1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minSwaps(tt.args.in0), "minSwaps(%v)", tt.args.in0)
		})
	}
}
