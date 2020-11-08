package slices

import (
	"fmt"
	"testing"
)

var (
	resultSlice []string
	resultInt   int

	s51  = "yyy"
	a100 = (func() []string {
		a := make([]string, 100)
		for i := 0; i < 100; i++ {
			if i == 50 {
				a[i] = s51
				continue
			}
			a[i] = "x" + fmt.Sprint(i)
		}
		return a
	})()
	aa100 = Split(a100, "")
)

func benchmarkMerge(b *testing.B, f func(...[]string) []string) {
	var a []string
	for i := 0; i < b.N; i++ {
		a = f(aa100...)
	}
	resultSlice = a
}

func BenchmarkMerge(b *testing.B) { benchmarkMerge(b, Merge) }

func benchmarkCount(b *testing.B, f func([]string, string) int) {
	var x int
	for i := 0; i < b.N; i++ {
		x = f(a100, s51)
	}
	resultInt = x
}

func BenchmarkCount(b *testing.B) { benchmarkCount(b, Count) }

func BenchmarkCount2(b *testing.B) {
	benchmarkCount(b,
		func(a []string, s string) int {
			var n int
			for i, j := 0, len(a)-1; i <= j; i, j = i+1, j-1 {
				if a[i] == s {
					n++
				}
				if i == j {
					break
				}
				if a[j] == s {
					n++
				}
			}
			return n
		})
}

func BenchmarkCount3(b *testing.B) {
	benchmarkCount(b,
		func(a []string, s string) int {
			n := 0
			for {
				i := Index(a, s)
				if i == -1 {
					return n
				}
				n++
				a = a[i+1:]
			}
		})
}
