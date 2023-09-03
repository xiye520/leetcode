package jianzhi_offer

import "testing"

func Test_sumNums(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"3",
			args{3},
			6,
		},
		{
			"9",
			args{9},
			45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNums(tt.args.n); got != tt.want {
				t.Errorf("sumNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
