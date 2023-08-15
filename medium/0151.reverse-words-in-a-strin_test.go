package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"the sky is blue",
			args{"the sky is blue"},
			"blue is sky the",
		},
		{
			"  hello world  ",
			args{"  hello world  "},
			"world hello",
		},
		{
			"a good   example",
			args{"a good   example"},
			"example good a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, reverseWords(tt.args.s), "reverseWords(%v)", tt.args.s)
		})
	}
}
