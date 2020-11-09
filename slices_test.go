// Copyright (c) 2019 srfrog - https://srfrog.me
// Use of this source code is governed by the license in the LICENSE file.

package slices

import (
	"reflect"
	"strings"
	"testing"
)

var (
	slc = []string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "", "elit", "donec", "tempus", "Lorem"}
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
		out := Index(slc, tc.in)
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
		{in: slc, s: "Lorem", out: 2},
		{in: slc, s: "srfrog", out: 0},
		{in: slc, s: "", out: 1},
		{in: nil, s: "Lorem", out: 0},
		{in: []string{"", "Lorem", ""}, s: "Lorem", out: 1},
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
		{in: slc[:5], f: func(s string) string { return s },
			out: []string{"Lorem", "ipsum", "dolor", "sit", "amet"}},
		{in: slc[:5], f: func(s string) string { return strings.Title(s) },
			out: []string{"Lorem", "Ipsum", "Dolor", "Sit", "Amet"}},
		{in: slc[:5], f: func(s string) string { return "XxX" + s },
			out: []string{"XxXLorem", "XxXipsum", "XxXdolor", "XxXsit", "XxXamet"}},
		{in: slc[:5], f: nil,
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
		{in: slc, f: func(s string) bool { return s == "Lorem" },
			out: []string{"ipsum", "dolor", "sit", "amet", "consectetur", "", "elit", "donec", "tempus"}},
		{in: slc, f: func(s string) bool { return s == "srfrog" },
			out: slc},
		{in: nil, f: func(s string) bool { return s == "srfrog" },
			out: []string{}},
		{in: slc, f: nil, out: slc},
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
		{in: slc, f: func(s string) bool { return s == "Lorem" },
			out: []string{"Lorem", "Lorem"}},
		{in: slc, f: func(s string) bool { return s == "srfrog" },
			out: []string{}},
		{in: nil, f: func(s string) bool { return s == "srfrog" },
			out: []string{}},
		{in: slc, f: nil, out: nil},
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
		{in: []string{}, out: len(slc)},
		{in: nil, out: len(slc)},
		{in: []string{"Lorem", "ipsum", "bacon"}, out: -1},
		{in: []string{"Lorem", "ipsum", "bacon", "ipsum"}, out: -2},
		{in: []string{"Florem", "ipsum"}, out: -2},
	}

	for _, tc := range tests {
		out := Compare(slc, tc.in)
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

func TestSplit(t *testing.T) {
	type args struct {
		a   []string
		sep string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{name: "empty slice",
			args: args{a: []string{}, sep: ""}, want: [][]string{}},
		{name: "nil",
			args: args{a: nil, sep: ""}, want: [][]string{}},
		{name: "sep empty",
			args: args{a: []string{"1", "2", "3"}, sep: ""},
			want: [][]string{{"1"}, {"2"}, {"3"}}},
		{name: "mismatch",
			args: args{a: []string{"1", "2", "3"}, sep: "horse"},
			want: [][]string{{"1", "2", "3"}}},
		{name: "match",
			args: args{a: []string{"law", "and", "order"}, sep: "and"},
			want: [][]string{{"law"}, {"order"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Split(tt.args.a, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitN(t *testing.T) {
	type args struct {
		a   []string
		sep string
		n   int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{name: "empty slice 0,-1",
			args: args{a: []string{}, sep: "", n: -1}, want: [][]string{}},
		{name: "empty slice 0,0",
			args: args{a: []string{}, sep: "", n: 0}, want: nil},
		{name: "empty slice 0,1",
			args: args{a: []string{}, sep: "", n: 1}, want: [][]string{}},
		{name: "nil",
			args: args{a: nil, sep: "", n: 0}, want: nil},
		{name: "sep empty",
			args: args{a: []string{"1", "2", "3"}, sep: "", n: 1},
			want: [][]string{{"1"}, {"2"}, {"3"}}},
		{name: "mismatch 0,-1",
			args: args{a: []string{"1", "2", "3"}, sep: "horse", n: -1},
			want: [][]string{{"1", "2", "3"}}},
		{name: "mismatch 0,1",
			args: args{a: []string{"1", "2", "3"}, sep: "horse", n: 1},
			want: [][]string{{"1", "2", "3"}}},
		{name: "match 1,1",
			args: args{a: []string{"law", "and", "order"}, sep: "and", n: 1},
			want: [][]string{{"law"}, {"order"}}},
		{name: "match 1,2",
			args: args{a: []string{"law", "and", "order"}, sep: "and", n: 2},
			want: [][]string{{"law"}, {"order"}}},
		{name: "match 3,1",
			args: args{a: []string{"Pig", "and", "Ale", "and", "Bar", "and", "Inn"}, sep: "and", n: 1},
			want: [][]string{{"Pig"}, {"Ale", "and", "Bar", "and", "Inn"}}},
		{name: "match 3,2",
			args: args{a: []string{"Pig", "and", "Ale", "and", "Bar", "and", "Inn"}, sep: "and", n: 2},
			want: [][]string{{"Pig"}, {"Ale"}, {"Bar", "and", "Inn"}}},
		{name: "match 3,3",
			args: args{a: []string{"Pig", "and", "Ale", "and", "Bar", "and", "Inn"}, sep: "and", n: 3},
			want: [][]string{{"Pig"}, {"Ale"}, {"Bar"}, {"Inn"}}},
		{name: "match 3,-1",
			args: args{a: []string{"Pig", "and", "Ale", "and", "Bar", "and", "Inn"}, sep: "and", n: -1},
			want: [][]string{{"Pig"}, {"Ale"}, {"Bar"}, {"Inn"}}},
		{name: "whys",
			args: args{a: []string{"why", "why", "why"}, sep: "why", n: 1},
			want: [][]string{{}, {"why", "why"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitN(tt.args.a, tt.args.sep, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChunk(t *testing.T) {
	type args struct {
		a    []string
		size int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{name: "nil", args: args{a: nil, size: 0}, want: nil},
		{name: "0,-1", args: args{a: []string{}, size: -1}, want: nil},
		{name: "0,0", args: args{a: []string{}, size: 0}, want: nil},
		{name: "0,1", args: args{a: []string{}, size: 1}, want: [][]string{}},
		{name: "2,1", args: args{a: []string{"1", "2"}, size: 1}, want: [][]string{{"1"}, {"2"}}},
		{name: "2,2", args: args{a: []string{"1", "2"}, size: 2}, want: [][]string{{"1", "2"}}},
		{name: "7,1",
			args: args{a: []string{"1", "2", "3", "4", "5", "6", "7"}, size: 1},
			want: [][]string{{"1"}, {"2"}, {"3"}, {"4"}, {"5"}, {"6"}, {"7"}}},
		{name: "7,2",
			args: args{a: []string{"1", "2", "3", "4", "5", "6", "7"}, size: 2},
			want: [][]string{{"1", "2"}, {"3", "4"}, {"5", "6"}, {"7"}}},
		{name: "7,3",
			args: args{a: []string{"1", "2", "3", "4", "5", "6", "7"}, size: 3},
			want: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7"}}},
		{name: "7,4",
			args: args{a: []string{"1", "2", "3", "4", "5", "6", "7"}, size: 4},
			want: [][]string{{"1", "2", "3", "4"}, {"5", "6", "7"}}},
		{name: "7,5",
			args: args{a: []string{"1", "2", "3", "4", "5", "6", "7"}, size: 5},
			want: [][]string{{"1", "2", "3", "4", "5"}, {"6", "7"}}},
		{name: "7,6",
			args: args{a: []string{"1", "2", "3", "4", "5", "6", "7"}, size: 6},
			want: [][]string{{"1", "2", "3", "4", "5", "6"}, {"7"}}},
		{name: "7,7",
			args: args{a: []string{"1", "2", "3", "4", "5", "6", "7"}, size: 7},
			want: [][]string{{"1", "2", "3", "4", "5", "6", "7"}}},
		{name: "7,8",
			args: args{a: []string{"1", "2", "3", "4", "5", "6", "7"}, size: 8},
			want: [][]string{{"1", "2", "3", "4", "5", "6", "7"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args.a, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "nil", args: args{a: nil, b: nil}, want: nil},
		{name: "none", args: args{a: []string{"1"}, b: []string{"2"}}, want: nil},
		{name: "one", args: args{a: []string{"1"}, b: []string{"1"}}, want: []string{"1"}},
		{name: "1,3",
			args: args{a: []string{"1", "2", "3"}, b: []string{"1", "4", "3"}},
			want: []string{"1", "3"}},
		{name: "3",
			args: args{a: []string{"1", "2", "3"}, b: []string{"3", "3", "3"}},
			want: []string{"3"}},
		{name: "3,3,3",
			args: args{a: []string{"3", "3", "3"}, b: []string{"1", "4", "3"}},
			want: []string{"3", "3", "3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	type args struct {
		ss [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "nil", args: args{ss: nil}, want: nil},
		{name: "n=0",
			args: args{ss: [][]string{{""}}},
			want: []string{""}},
		{name: "n=1",
			args: args{ss: [][]string{{"1"}}},
			want: []string{"1"}},
		{name: "n=2",
			args: args{ss: [][]string{{"1"}, {"2", "2"}}},
			want: []string{"1", "2", "2"}},
		{name: "n=3",
			args: args{ss: [][]string{{"1"}, {"2", "2"}, {"3", "3", "3"}}},
			want: []string{"1", "2", "2", "3", "3", "3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.ss...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffle(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "nil", args: args{a: nil}, want: nil},
		{name: "n=0", args: args{a: []string{}}, want: []string{}},
		{name: "n=1", args: args{a: []string{"1"}}, want: []string{"1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Shuffle(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertAt(t *testing.T) {
	type args struct {
		a      []string
		idx    int
		values []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "nil", args: args{a: nil, idx: 0, values: nil}, want: nil},
		{name: "empty", args: args{a: []string{}, idx: 0, values: []string{}}, want: []string{}},
		{name: "a=0,v=1,idx=0",
			args: args{a: []string{}, idx: 0, values: []string{"1"}},
			want: []string{"1"}},
		{name: "a=3,v=1,idx=0",
			args: args{a: []string{"a", "b", "c"}, idx: 0, values: []string{"1"}},
			want: []string{"1", "a", "b", "c"}},
		{name: "a=3,v=1,idx=1",
			args: args{a: []string{"a", "b", "c"}, idx: 1, values: []string{"1"}},
			want: []string{"a", "1", "b", "c"}},
		{name: "a=3,v=1,idx=2",
			args: args{a: []string{"a", "b", "c"}, idx: 2, values: []string{"1"}},
			want: []string{"a", "b", "1", "c"}},
		{name: "a=3,v=1,idx=3",
			args: args{a: []string{"a", "b", "c"}, idx: 3, values: []string{"1"}},
			want: []string{"a", "b", "c", "1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertAt(tt.args.a, tt.args.idx, tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPop(t *testing.T) {
	slc := []string{"1", "2", "3"}

	type args struct {
		a *[]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "empty", args: args{a: &([]string{})}, want: ""},
		{name: "pop-1", args: args{a: &slc}, want: "3"},
		{name: "pop-2", args: args{a: &slc}, want: "2"},
		{name: "pop-3", args: args{a: &slc}, want: "1"},
		{name: "pop-4", args: args{a: &slc}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pop(tt.args.a); got != tt.want {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPush(t *testing.T) {
	var slc []string

	type args struct {
		a *[]string
		s []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "empty", args: args{a: &([]string{}), s: []string{}}, want: 0},
		{name: "push-1", args: args{a: &slc, s: []string{"1"}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Push(tt.args.a, tt.args.s...); got != tt.want {
				t.Errorf("Push() = %v, want %v", got, tt.want)
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
		{name: "nil", args: args{a: nil, b: nil}, want: true},
		{name: "equal", args: args{a: []string{"1"}, b: []string{"1"}}, want: true},
		{name: "a-longer", args: args{a: []string{"1", "2"}, b: []string{"1"}}, want: false},
		{name: "b-longer", args: args{a: []string{"1"}, b: []string{"1", "2"}}, want: false},
		{name: "a-empty", args: args{a: []string{}, b: []string{"1", "2"}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice(t *testing.T) {
	slc := []string{"a", "b", "c", "d", "e"}

	type args struct {
		a      []string
		offset int
		length int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "nil", args: args{a: nil, offset: -1, length: 0}, want: nil},
		{name: "a=1,offset=-1,length=0",
			args: args{a: []string{"1"}, offset: -1, length: 0},
			want: []string{"1"}},
		{name: "a=5,offset=2,length=0",
			args: args{a: slc, offset: 2, length: 0},
			want: []string{"c", "d", "e"}},
		{name: "a=5,offset=-2,length=1",
			args: args{a: slc, offset: -2, length: 1},
			want: []string{"d"}},
		{name: "a=5,offset=0,length=3",
			args: args{a: slc, offset: 0, length: 3},
			want: []string{"a", "b", "c"}},
		{name: "a=5,offset=1,length=2",
			args: args{a: slc, offset: 1, length: 2},
			want: []string{"b", "c"}},
		{name: "a=5,offset=-3,length=-4",
			args: args{a: slc, offset: -3, length: -4},
			want: nil},
	}
	// slc = []string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "", "elit", "donec", "tempus", "Lorem"}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slice(tt.args.a, tt.args.offset, tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplice(t *testing.T) {
	slc := []string{"a", "b", "c"}

	type args struct {
		a      []string
		offset int
		length int
		b      []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "nil", args: args{a: nil, b: nil}, want: nil},
		{name: "a=0,b=1,offset=1,length=0",
			args: args{a: nil, offset: 1, length: 0, b: []string{"1"}},
			want: nil},
		{name: "a=3,b=2,offset=-1,length=1",
			args: args{a: slc, offset: -1, length: 1, b: []string{"1", "2"}},
			want: []string{"a", "b", "1", "2"}},
		{name: "a=3,b=1,offset=1,length=max",
			args: args{a: slc, offset: 1, length: len(slc), b: []string{"1"}},
			want: []string{"a", "1"}},
		{name: "a=3,b=0,offset=1,length=1",
			args: args{a: slc, offset: 1, length: 1, b: nil},
			want: []string{"a", "c"}},
		{name: "a=3,b=0,offset=-1,length=1",
			args: args{a: slc, offset: -1, length: 1, b: nil},
			want: []string{"a", "c"}},
		{name: "a=10,b=3,offset=3,length=3",
			args: args{a: Repeat("x", 10), offset: 3, length: 3, b: []string{"1", "2", "3"}},
			want: []string{"x", "x", "x", "1", "2", "3", "x", "x", "x", "x"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Splice(tt.args.a, tt.args.offset, tt.args.length, tt.args.b...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Splice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "nil", args: args{a: nil}, want: nil},
		{name: "diff=3",
			args: args{a: []string{"1", "2", "3"}},
			want: []string{"1", "2", "3"}},
		{name: "diff=2",
			args: args{a: []string{"1", "1", "3"}},
			want: []string{"1", "3"}},
		{name: "diff=1",
			args: args{a: []string{"1", "1", "1"}},
			want: []string{"1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
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
		{name: "a=nil,b=nil", args: args{a: nil, b: nil}, want: -1},
		{name: "a=nil,b=3",
			args: args{a: nil, b: []string{"1", "2", "3"}},
			want: -1},
		{name: "a=1,b=3,-1",
			args: args{a: []string{"x"}, b: []string{"1", "2", "3"}},
			want: -1},
		{name: "a=1,b=3,0",
			args: args{a: []string{"2"}, b: []string{"1", "2", "3"}},
			want: 0},
		{name: "a=3,b=3,0",
			args: args{a: []string{"3", "2", "1"}, b: []string{"1", "2", "3"}},
			want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexAny(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IndexAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastIndexAny(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "a=nil,b=nil", args: args{a: nil, b: nil}, want: -1},
		{name: "a=nil,b=3",
			args: args{a: nil, b: []string{"1", "2", "3"}},
			want: -1},
		{name: "a=1,b=3,-1",
			args: args{a: []string{"x"}, b: []string{"1", "2", "3"}},
			want: -1},
		{name: "a=1,b=3,0",
			args: args{a: []string{"2"}, b: []string{"1", "2", "3"}},
			want: 0},
		{name: "a=3,b=3,0",
			args: args{a: []string{"3", "2", "1"}, b: []string{"1", "2", "3"}},
			want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastIndexAny(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("LastIndexAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepeat(t *testing.T) {
	type args struct {
		s     string
		count int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "empty", args: args{s: "", count: 0}, want: []string{}},
		{name: "x1",
			args: args{s: "x", count: 1},
			want: []string{"x"}},
		{name: "x5",
			args: args{s: "x", count: 5},
			want: []string{"x", "x", "x", "x", "x"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.args.s, tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	type args struct {
		a      []string
		substr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "nil", args: args{a: nil, substr: ""}, want: -1},
		{name: "a=3,empty",
			args: args{a: []string{"a", "", "c"}, substr: ""},
			want: 0},
		{name: "a=3,space",
			args: args{a: []string{"a", " ", "c"}, substr: " "},
			want: 1},
		{name: "a=3,apple",
			args: args{a: []string{"orange", "quenepas", "crabapple", "kiwi"}, substr: "apple"},
			want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.a, tt.args.substr); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	type args struct {
		a []string
		f func(string, int, string) string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "single",
			args: args{
				a: []string{"1"},
				f: func(acc string, i int, v string) string { return "" }},
			want: "1"},
		{name: "max",
			args: args{
				a: []string{"a", "b", "a", "b", "c", "e", "e", "c", "d", "d", "d", "d"},
				f: func(acc string, i int, v string) string {
					if v > acc {
						return v
					}
					return acc
				}},
			want: "e"},
		{name: "join",
			args: args{
				a: []string{"1", "2", "3", "4"},
				f: func(acc string, i int, v string) string {
					if acc != "" {
						acc = acc + ", "
					}
					return acc + v
				}},
			want: "1, 2, 3, 4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.a, tt.args.f); got != tt.want {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
