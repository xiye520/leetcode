package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{3, 7},
			28,
		},
		{
			"",
			args{3, 2},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, uniquePaths(tt.args.m, tt.args.n), "uniquePaths2(%v, %v)", tt.args.m, tt.args.n)
		})
	}
}

func Test_uniquePaths2(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{3, 7},
			28,
		},
		{
			"",
			args{3, 2},
			3,
		},
		{
			"",
			args{1, 1},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, uniquePaths2(tt.args.m, tt.args.n), "uniquePaths2(%v, %v)", tt.args.m, tt.args.n)
		})
	}
}
