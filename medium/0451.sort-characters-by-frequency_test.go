package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_frequencySort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"eert",
			args{"tree"},
			"eetr",
		},
		{
			"cccaaa",
			args{"cccaaa"},
			"cccaaa",
		},
		{
			"Aabb",
			args{"bbAa"},
			"bbAa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, frequencySort(tt.args.s), "frequencySort(%v)", tt.args.s)
		})
	}
}
