package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"ab",
			args{"ab", "eidbaooo"},
			true,
		},
		{
			"",
			args{"ab", "eidboaoo"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, checkInclusion(tt.args.s1, tt.args.s2), "checkInclusion(%v, %v)", tt.args.s1, tt.args.s2)
		})
	}
}
