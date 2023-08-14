package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isBalanceStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"success",
			args{"aba"},
			true,
		},
		{
			"abab not balance",
			args{"abab"},
			false,
		},
		{
			"success",
			args{"ababa"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isBalanceStr(tt.args.s), "isBalanceStr(%v)", tt.args.s)
		})
	}
}

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{"aba"},
			"aba",
		},
		{
			"2",
			args{"babad"},
			"bab",
		},
		{
			"3",
			args{"cbbd"},
			"bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, longestPalindrome(tt.args.s), "longestPalindrome(%v)", tt.args.s)
		})
	}
}
