package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"b",
			args{"a", "b"},
			false,
		},
		{
			"ab",
			args{"aa", "ab"},
			false,
		},
		{
			"aab",
			args{"aa", "aab"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, canConstruct(tt.args.ransomNote, tt.args.magazine), "canConstruct(%v, %v)", tt.args.ransomNote, tt.args.magazine)
		})
	}
}
