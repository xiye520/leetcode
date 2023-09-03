package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"cbaebabacd",
			args{"cbaebabacd", "abc"},
			[]int{0, 6},
		},
		{
			"abab",
			args{"abab", "ab"},
			[]int{0, 1, 2},
		},
		{
			"aaaaaaaaaa",
			args{"aaaaaaaaaa", "aaaaaaaaaaaaa"},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findAnagrams(tt.args.s, tt.args.p), "findAnagrams(%v, %v)", tt.args.s, tt.args.p)
		})
	}
}
