package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxScore1(t *testing.T) {
	type args struct {
		cardPoints []int
		k          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]int{1, 2, 3, 4, 5, 6, 1}, 3},
			12,
		},
		{
			"",
			args{[]int{2, 2, 2}, 2},
			4,
		},
		{
			"",
			args{[]int{9, 7, 7, 9, 7, 7, 9}, 7},
			55,
		},
		{
			"",
			args{[]int{1, 1000, 1}, 1},
			1,
		},
		{
			"",
			args{[]int{1, 79, 80, 1, 1, 1, 200, 1}, 3},
			202,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxScore(tt.args.cardPoints, tt.args.k), "maxScore1(%v, %v)", tt.args.cardPoints, tt.args.k)
		})
	}
}
