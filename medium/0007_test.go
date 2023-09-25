package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{123},
			321,
		},
		{
			"",
			args{-123},
			-321,
		},
		{
			"",
			args{120},
			21,
		},
		{
			"",
			args{0},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, reverse(tt.args.x), "reverse(%v)", tt.args.x)
		})
	}
}
