// Copyright (c) 2019 srfrog - https://srfrog.me
// Use of this source code is governed by the license in the LICENSE file.

package slices

import (
	"strings"
	"testing"
)

var (
	arr  = []string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "", "elit", "donec", "tempus", "Lorem"}
	arr2 = []string{"*L*", "*i*", "*d*", "*s*", "*a*", "*c*", "*a*", "*e*", "*d*", "*t*", "*L*"}
)

func TestIndex(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{in: "Lorem", out: 0},
		{in: "dolor", out: 2},
		{in: "horse", out: -1},
		{in: "", out: 6},
	}

	for _, tc := range tests {
		out := Index(arr, tc.in)
		if out != tc.out {
			t.Errorf("expecting %v got %v", tc.out, out)
		}
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		in  []string
		s   string
		out int
	}{
		{in: arr, s: "Lorem", out: 2},
		{in: arr, s: "srfrog", out: 0},
		{in: arr, s: "", out: 1},
		{in: nil, s: "Lorem", out: 0},
	}

	for _, tc := range tests {
		out := Count(tc.in, tc.s)
		if out != tc.out {
			t.Errorf("expecting %v got %v", tc.out, out)
		}
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		in  []string
		f   func(string) string
		out []string
	}{
		{in: arr[:5], f: func(s string) string { return s },
			out: []string{"Lorem", "ipsum", "dolor", "sit", "amet"}},
		{in: arr[:5], f: func(s string) string { return strings.Title(s) },
			out: []string{"Lorem", "Ipsum", "Dolor", "Sit", "Amet"}},
		{in: arr[:5], f: func(s string) string { return "XxX" + s },
			out: []string{"XxXLorem", "XxXipsum", "XxXdolor", "XxXsit", "XxXamet"}},
		{in: arr[:5], f: nil,
			out: []string{"Lorem", "ipsum", "dolor", "sit", "amet"}},
		{in: nil, f: strings.Title,
			out: []string{}},
	}

	for _, tc := range tests {
		out := Map(tc.f, tc.in)
		if !Equal(tc.out, out) {
			t.Errorf("expecting %v got %v", tc.out, out)
		}
	}
}

func TestTrimFunc(t *testing.T) {
	tests := []struct {
		in  []string
		f   func(string) bool
		out []string
	}{
		{in: arr[:5], f: func(s string) bool { return s == "Lorem" },
			out: []string{"ipsum", "dolor", "sit", "amet"}},
		{in: arr[:5], f: func(s string) bool { return s == "srfrog" },
			out: []string{"Lorem", "ipsum", "dolor", "sit", "amet"}},
		{in: nil, f: func(s string) bool { return s == "srfrog" },
			out: []string{}},
		{in: arr[:5], f: nil,
			out: []string{"Lorem", "ipsum", "dolor", "sit", "amet"}},
	}

	for _, tc := range tests {
		out := TrimFunc(tc.in, tc.f)
		if !Equal(out, tc.out) {
			t.Errorf("expecting %v got %v", tc.out, out)
		}
	}
}

func TestFilterFunc(t *testing.T) {
	tests := []struct {
		in  []string
		f   func(string) bool
		out []string
	}{
		{in: arr[:5], f: func(s string) bool { return s == "Lorem" },
			out: []string{"Lorem"}},
		{in: arr[:5], f: func(s string) bool { return s == "srfrog" },
			out: []string{}},
		{in: nil, f: func(s string) bool { return s == "srfrog" },
			out: []string{}},
		{in: arr[:5], f: nil,
			out: []string{"Lorem", "ipsum", "dolor", "sit", "amet"}},
	}

	for _, tc := range tests {
		out := FilterFunc(tc.in, tc.f)
		if !Equal(out, tc.out) {
			t.Errorf("expecting %v got %v", tc.out, out)
		}
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		in  []string
		out int
	}{
		{in: []string{"Lorem"}, out: 0},
		{in: []string{"Lorem", "ipsum"}, out: 0},
		{in: []string{}, out: len(arr)},
		{in: nil, out: len(arr)},
		{in: []string{"Lorem", "ipsum", "bacon"}, out: -1},
		{in: []string{"Lorem", "ipsum", "bacon", "ipsum"}, out: -2},
		{in: []string{"Florem", "ipsum"}, out: -2},
	}

	for _, tc := range tests {
		out := Compare(arr, tc.in)
		if out != tc.out {
			t.Errorf("expecting %v got %v", tc.out, out)
		}
	}
}

func TestReplace(t *testing.T) {
	a := []string{"Lorem", "", "Lorem"}
	tests := []struct {
		old, new string
		n        int
		out      []string
	}{
		{old: "Lorem", new: "Florem", n: 1,
			out: []string{"Florem", "", "Lorem"}},
		{old: "Lorem", new: "Florem", n: 0,
			out: []string{"Lorem", "", "Lorem"}},
		{old: "Lorem", new: "Florem", n: 2,
			out: []string{"Florem", "", "Florem"}},
		{old: "Lorem", new: "Florem", n: -1,
			out: []string{"Florem", "", "Florem"}},
		{old: "Lorem", new: "Florem", n: 100,
			out: []string{"Florem", "", "Florem"}},
		{old: "Lorem", new: "Florem", n: -100,
			out: []string{"Florem", "", "Florem"}},
		{old: "Lorem", new: "Lorem", n: 1,
			out: []string{"Lorem", "", "Lorem"}},
		{old: "", new: "Lorem", n: 1,
			out: []string{"Lorem", "Lorem", "Lorem"}},
	}

	for _, tc := range tests {
		out := Replace(a, tc.old, tc.new, tc.n)
		if !Equal(tc.out, out) {
			t.Errorf("expecting %v got %v", tc.out, out)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		in, out []string
	}{
		{in: nil, out: nil},
		{in: []string{"Lorem", "ipsum"},
			out: []string{"ipsum", "Lorem"}},
		{in: []string{"Lorem", "ipsum", "dolor", "sit", "amet"},
			out: []string{"amet", "sit", "dolor", "ipsum", "Lorem"}},
	}

	for _, tc := range tests {
		out := Reverse(tc.in)
		if !Equal(tc.out, out) {
			t.Errorf("expecting %v got %v", tc.out, out)
		}
	}
}

func TestIndexAny(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexAny(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IndexAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexFunc(t *testing.T) {
	type args struct {
		a []string
		f func(string) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexFunc(tt.args.a, tt.args.f); got != tt.want {
				t.Errorf("IndexFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	type args struct {
		a []string
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.a, tt.args.s); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualFold(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EqualFold(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("EqualFold() = %v, want %v", got, tt.want)
			}
		})
	}
}
