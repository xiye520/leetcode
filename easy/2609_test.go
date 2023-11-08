package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findTheLongestBalancedSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{"01000111"},
			6,
		},
		{
			"",
			args{"00111"},
			4,
		},
		{
			"",
			args{"111"},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findTheLongestBalancedSubstring(tt.args.s), "findTheLongestBalancedSubstring(%v)", tt.args.s)
		})
	}
}
