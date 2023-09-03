package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_equalSubstring(t *testing.T) {
	type args struct {
		s       string
		t       string
		maxCost int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{"abcd", "bcdf", 3},
			3,
		},
		{
			"",
			args{"abcd", "cdef", 3},
			1,
		},
		{
			"",
			args{"abcd", "acde", 0},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, equalSubstring(tt.args.s, tt.args.t, tt.args.maxCost), "equalSubstring(%v, %v, %v)", tt.args.s, tt.args.t, tt.args.maxCost)
		})
	}
}

func Test_equalSubstring2(t *testing.T) {
	type args struct {
		s       string
		t       string
		maxCost int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{"abcd", "bcdf", 3},
			3,
		},
		{
			"abcd",
			args{"abcd", "cdef", 3},
			1,
		},
		{
			"",
			args{"abcd", "acde", 0},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, equalSubstring2(tt.args.s, tt.args.t, tt.args.maxCost), "equalSubstring(%v, %v, %v)", tt.args.s, tt.args.t, tt.args.maxCost)
		})
	}
}
