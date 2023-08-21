package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"abba",
			args{"abba", "dog cat cat dog"},
			true,
		},
		{
			"dog cat cat fish",
			args{"abba", "dog cat cat fish"},
			false,
		},
		{
			"dog cat cat dog",
			args{"aaaa", "dog cat cat dog"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, wordPattern(tt.args.pattern, tt.args.s), "wordPattern(%v, %v)", tt.args.pattern, tt.args.s)
		})
	}
}
