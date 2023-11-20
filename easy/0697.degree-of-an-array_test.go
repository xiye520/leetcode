package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findShortestSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			"",
			args{[]int{1, 2, 2, 3, 1}},
			2,
		},
		{
			"",
			args{[]int{1, 2, 2, 3, 1, 4, 2}},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantAns, findShortestSubArray(tt.args.nums), "findShortestSubArray(%v)", tt.args.nums)
		})
	}
}
