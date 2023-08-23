package main

import "testing"

func Test_getCompareRoute(t *testing.T) {
	type args struct {
		routes []string
		url    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"/a*/cde",
			args{[]string{"/a*/cde", "/af/cde"}, "/ab/cde"},
			"/a*/cde",
		},
		{
			"/ab/cde",
			args{[]string{"/a*/cde", "/af/cde", "/ab/cde"}, "/ab/cde"},
			"/ab/cde",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCompareRoute(tt.args.routes, tt.args.url); got != tt.want {
				t.Errorf("getCompareRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxCompareCount(t *testing.T) {
	type args struct {
		route string
		url   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{"/a*/cde", "/ab/cdef"},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCompareCount(tt.args.route, tt.args.url); got != tt.want {
				t.Errorf("maxCompareCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMaxPrefix(t *testing.T) {
	type args struct {
		route string
		url   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{"abc", "abcde"},
			3,
		},
		{
			"",
			args{"abc", "abdcde"},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxPrefix(tt.args.route, tt.args.url); got != tt.want {
				t.Errorf("getMaxPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
