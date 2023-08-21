package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLongestSubstring6(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"abcabcbb",
			args{"abcabcbb"},
			3,
		},
		{
			"bbbbb",
			args{"bbbbb"},
			1,
		},
		{
			"pwwkew",
			args{"pwwkew"},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, lengthOfLongestSubstring6(tt.args.s), "lengthOfLongestSubstring6(%v)", tt.args.s)
		})
	}
}
