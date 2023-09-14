package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"22",
			args{
				[][]int{
					{1, 3, 1, 6},
					{1, 5, 1, 3},
					{4, 2, 1, 2}},
			},
			9,
		},
		{
			"22",
			args{
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minPathSum(tt.args.grid), "minPathSum(%v)", tt.args.grid)
		})
	}
}

func Test_minPathSum2(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"22",
			args{
				[][]int{
					{1, 3, 1, 6},
					{1, 5, 1, 3},
					{4, 2, 1, 2}},
			},
			9,
		},
		{
			"22",
			args{
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minPathSum4(tt.args.grid), "minPathSum2(%v)", tt.args.grid)
		})
	}
}
