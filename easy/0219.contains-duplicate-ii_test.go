package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_containsNearbyDuplicate2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1,2,3,1",
			args{[]int{1, 2, 3, 1}, 3},
			true,
		},
		{
			"1,0,1,1",
			args{[]int{1, 0, 1, 1}, 1},
			true,
		},
		{
			"1,2,3,1,2,3",
			args{[]int{1, 2, 3, 1, 2, 3}, 2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, containsNearbyDuplicate2(tt.args.nums, tt.args.k), "containsNearbyDuplicate2(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
