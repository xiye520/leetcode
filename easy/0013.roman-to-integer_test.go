package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "III-3",
			args: args{"III"},
			want: 3,
		},
		{
			name: "IV-4",
			args: args{"IV"},
			want: 4,
		},
		{
			name: "LVIII-58",
			args: args{"LVIII"},
			want: 58,
		},
		{
			name: "MCMXCIV-1994",
			args: args{"MCMXCIV"},
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, romanToInt(tt.args.s), "romanToInt(%v)", tt.args.s)
		})
	}
}
