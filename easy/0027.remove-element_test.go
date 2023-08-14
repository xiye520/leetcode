package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeElement2(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"success",
			args{[]int{3, 2, 2, 3}, 3},
			2,
		},
		{
			"success",
			args{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
			5,
		},
		{
			"success",
			args{[]int{1}, 1},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, removeElement2(tt.args.nums, tt.args.val), "removeElement2(%v, %v)", tt.args.nums, tt.args.val)
		})
	}
}
