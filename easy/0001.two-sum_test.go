package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 1.两数之和
func TestTwoSum3(t *testing.T) {
	target := 9
	nums := []int{2, 7, 11, 15}

	//indexs := twoSum2(nums, target)
	indexs := twoSum3(nums, target)

	if len(nums) < 2 || nums[indexs[0]]+nums[indexs[1]] != target {
		t.Fatal("fails....")
	}
}

func Test_twoSum5(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"2,7,11,15",
			args{[]int{2, 7, 11, 15}, 9},
			[]int{0, 1},
		},
		{
			"3,2,4",
			args{[]int{3, 2, 4}, 6},
			[]int{1, 2},
		},
		{
			"3,3",
			args{[]int{3, 3}, 6},
			[]int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, twoSum5(tt.args.nums, tt.args.target), "twoSum5(%v, %v)", tt.args.nums, tt.args.target)
		})
	}
}
