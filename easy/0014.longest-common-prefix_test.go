package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"success",
			args{[]string{"flower", "flow", "flight"}},
			"fl",
		},
		{
			"success",
			args{[]string{"dog", "racecar", "car"}},
			"",
		},
		{
			"success",
			args{[]string{"ab", "a"}},
			"a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, longestCommonPrefix2(tt.args.strs), "longestCommonPrefix(%v)", tt.args.strs)
		})
	}
}
