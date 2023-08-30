package jianzhi_offer

import "testing"

func Test_numWays(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{1},
			1,
		},
		{
			"",
			args{0},
			1,
		},
		{
			"",
			args{2},
			2,
		},
		{
			"",
			args{7},
			21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWays(tt.args.n); got != tt.want {
				t.Errorf("numWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
