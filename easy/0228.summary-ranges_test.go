package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"0,1,2,4,5,7",
			args{[]int{0, 1, 2, 4, 5, 7}},
			[]string{"0->2", "4->5", "7"},
		},
		{
			"0,2,3,4,6,8,9",
			args{[]int{0, 2, 3, 4, 6, 8, 9}},
			[]string{"0", "2->4", "6", "8->9"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, summaryRanges(tt.args.nums), "summaryRanges(%v)", tt.args.nums)
		})
	}
}
