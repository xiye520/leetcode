package hard

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"",
			args{"aa", "a"},
			false,
		},
		{
			"",
			args{"aa", "a*"},
			true,
		},
		{
			"",
			args{"ab", ".*"},
			true,
		},
		{
			"",
			args{"aab", "c*a*b"},
			true,
		},
		{
			"",
			args{"mississippi", "mis*is*p*."},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
