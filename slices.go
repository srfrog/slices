// Copyright (c) 2019 srfrog - https://srfrog.me
// Use of this source code is governed by the license in the LICENSE file.

// Package slices is a collection of functions to manipulate string slices.
// Some functions were adapted from the strings package to work with string slices, other
// were ported from PHP 'array_*' function equivalents.
package slices

import (
	"math/rand"
	"strings"
)

// streq just compares 2 strings, used as base comparison.
func streq(s1, s2 string) bool { return s1 == s2 }

// Compare returns an integer comparing two string slices lexicographically.
// The result will be 0 if a==b, or that a has all values of b.
// The result will be -1 if a < b, or a is shorter than b.
// The result will be +1 if a > b, or a is longer than b.
// A nil argument is equivalent to an empty slice.
func Compare(a, b []string) int {
	return compareFunc(a, b, streq)
}

func compareFunc(a, b []string, f func(string, string) bool) int {
	var i int

	m, n := len(a), len(b)
	switch {
	case m == 0:
		return -n
	case n == 0:
		return m
	case m > n:
		m = n
	}

	for i = 0; i < m; i++ {
		if !f(a[i], b[i]) {
			break
		}
	}

	return i - n
}

// Contains returns true if 's' is in 'a', false otherwise
func Contains(a []string, s string) bool {
	return Index(a, s) >= 0
}

// ContainsAny returns true if any value in 'b' is in 'a', false otherwise
func ContainsAny(a, b []string) bool {
	return IndexAny(a, b) >= 0
}

// ContainsPrefix returns true if any entry in 'a' has prefix 'prefix', false otherwise
func ContainsPrefix(a []string, prefix string) bool {
	return indexFunc(a, prefix, strings.HasPrefix) >= 0
}

// ContainsSuffix returns true if any entry in 'a' has suffix 'suffix', false otherwise
func ContainsSuffix(a []string, suffix string) bool {
	return indexFunc(a, suffix, strings.HasSuffix) >= 0
}

// Count returns the number of occurrences of 's' in 'a'
func Count(a []string, s string) int {
	var n int
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		if a[i] == s {
			n++
		}
		if a[j] == s {
			n++
		}
	}
	return n
}

// Diff returns a slice with all the entries of 'a' that are not found in 'b'
func Diff(a, b []string) []string {
	return diffFunc(a, b, func(a []string, s string) bool { return !Contains(a, s) })
}

// diffFunc compares the elements of 'a' with those of 'b' using 'f' comparison function (bool).
// It returns a slice of the elements in 'a' that are not found in 'b' which f() == true.
func diffFunc(a, b []string, f func(a []string, s string) bool) []string {
	var c []string
	for i := range a {
		if f(b, a[i]) {
			c = append(c, a[i])
		}
	}
	return c
}

// Equal returns a boolean reporting whether a and b are the same length and contain the
// same values, when compared lexicographically.
func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	return Compare(a, b) == 0
}

// EqualFold returns a boolean reporting whether a and b
// are the same length and their values are equal under Unicode case-folding.
func EqualFold(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	return compareFunc(a, b, strings.EqualFold) == 0
}

// Fill is an alias of Repeat (PHP)
func Fill(n int, s string) []string {
	return Repeat(s, n)
}

// Filter returns a slice with all the entries of 'a' that match string 's'
func Filter(a []string, s string) []string {
	return filterFunc(a, s, streq)
}

// filterFunc runs through 'a' and compares each element with 's' using 'f' comparison
// function (bool). It returns a slice of the elements in 'a' where 'f' returns true.
func filterFunc(a []string, s string, f func(string, string) bool) []string {
	// complement trimFunc comparison
	return trimFunc(a, s, func(v string, s string) bool { return !f(v, s) })
}

// FilterFunc returns a slice with all the entries of 'a' that don't match string 's' that
// satisfy f(s). If func f returns true, the value will be filtered from 'a'.
func FilterFunc(a []string, f func(string) bool) []string {
	var b []string
	if f == nil {
		return a
	}
	for i := range a {
		if f(a[i]) {
			b = append(b, a[i])
		}
	}
	return b
}

// FilterPrefix returns a slice with all the entries of 'a' that have prefix 'prefix'
func FilterPrefix(a []string, prefix string) []string {
	return filterFunc(a, prefix, strings.HasPrefix)
}

// FilterSuffix returns a slice with all the entries of 'a' that have suffix 'suffix'
func FilterSuffix(a []string, suffix string) []string {
	return filterFunc(a, suffix, strings.HasSuffix)
}

// Index returns the index of the first instance of 's' in 'a', or -1 if not found
func Index(a []string, s string) int {
	return indexFunc(a, s, streq)
}

// IndexAny returns the index of the first instance of 'b' in 'a', or -1 if not found
func IndexAny(a, b []string) int {
	return fanOutFunc(a, b, Index)
}

// indexFunc runs through 'a' and compares each element with 's' using 'f' function.
// It returns the index of the first occurrence of 's' in 'a', or -1 if not found.
func indexFunc(a []string, s string, f func(string, string) bool) int {
	for i := range a {
		if f(a[i], s) {
			return i
		}
	}
	return -1
}

// IndexFunc returns the index into a of the first string satifying f(s), or -1 if none do.
func IndexFunc(a []string, f func(string) bool) int {
	for i := range a {
		if f(a[i]) {
			return i
		}
	}
	return -1
}

// Intersect returns a slice with all the entries of 'a' that are found in 'b'
func Intersect(a, b []string) []string {
	// complement diffFunc comparison
	return diffFunc(a, b, Contains)
}

// lastIndexFunc runs through 'a' and compares each element with 's' using 'f' function.
// It returns the index of the last occurrence of 's' in 'a', or -1 if not found.
func lastIndexFunc(a []string, s string, f func(string, string) bool) int {
	for i := len(a) - 1; i >= 0; i-- {
		if f(a[i], s) {
			return i
		}
	}
	return -1
}

// LastIndex returns the index of the last instance of 's' in 'a', or -1 if not found
func LastIndex(a []string, s string) int {
	return lastIndexFunc(a, s, streq)
}

// LastIndexAny returns the index of the last instance of 'b' in 'a', or -1 if not found
func LastIndexAny(a, b []string) int {
	return fanOutFunc(a, b, LastIndex)
}

// LastIndexFunc returns the index into a of the last string satifying f(s), or -1 if none do.
func LastIndexFunc(a []string, f func(string) bool) int {
	for i := len(a) - 1; i >= 0; i-- {
		if f(a[i]) {
			return i
		}
	}
	return -1
}

// LastSearch returns the index of the last entry containing the substring 's' in 'a',
// or -1 if not found
func LastSearch(a []string, s string) int {
	return lastIndexFunc(a, s, strings.Contains)
}

// Map returns a new slice with the function 'mapping' applied to each element of 'a'
func Map(mapping func(string) string, a []string) []string {
	if mapping == nil {
		return a
	}

	b := make([]string, len(a))
	for i := range a {
		b[i] = mapping(a[i])
	}
	return b
}

// Pop removes the last element in '*a' and returns it, shortening the array by one.
// If '*a' is empty returns empty string "".
// Note that this function will change the array pointed by 'a'.
func Pop(a *[]string) string {
	var s string
	if m := len(*a); m > 0 {
		s, *a = (*a)[m-1], (*a)[:m-1]
	}
	return s
}

// Push appends one or more elements to '*a' and returns the number of entries.
// Note that this function will change the array pointed by 'a'.
func Push(a *[]string, s ...string) int {
	if s != nil {
		*a = append(*a, s...)
	}
	return len(*a)
}

// Repeat returns a slice consisting of 'n' copies of 's'
func Repeat(s string, n int) []string {
	a := make([]string, n)
	for i := range a {
		a[i] = s
	}
	return a
}

// Replace returns a copy of the slice 'a' with the first 'n' instances of old replaced by new.
// If n < 0, there is no limit on the number of replacements.
func Replace(a []string, old, new string, n int) []string {
	m := len(a)
	if old == new || m == 0 || n == 0 {
		return a
	}

	if m := Count(a, old); m == 0 {
		return a
	} else if n < 0 || m < n {
		n = m
	}

	t := append(a[:0:0], a...)
	for i := 0; i < m; i++ {
		if n == 0 {
			break
		}
		if t[i] == old {
			t[i] = new
			n--
		}
	}
	return t
}

// ReplaceAll returns a copy of the slice 'a' with all instances of old replaced by new.
func ReplaceAll(a []string, old, new string) []string {
	return Replace(a, old, new, -1)
}

// Rand returns a new slice with 'n' number of random entries of 'a'.
// Note: You may want initialize the rand seed once in your program.
//
//    rand.Seed(time.Now().UnixNano())
//
func Rand(a []string, n int) []string {
	b := make([]string, n)
	if m := len(a); m > 0 {
		for i := 0; i < n; i++ {
			b[i] = a[rand.Intn(m)]
		}
	}
	return b
}

// Reverse returns a slice of the reverse index order elements of 'a'.
func Reverse(a []string) []string {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

// Search returns the index of the first entry containing the substring 's' in 'a',
// or -1 if not found
func Search(a []string, s string) int {
	return indexFunc(a, s, strings.Contains)
}

// Shift shifts the first element of '*a' and returns it, shortening the array by one.
// If '*a' is empty returns empty string "".
// Note that this function will change the array pointed by 'a'.
func Shift(a *[]string) string {
	var s string
	if m := len(*a); m > 0 {
		s, *a = (*a)[0], (*a)[1:]
	}
	return s
}

// Shuffle returns a slice with randomized order of elements in 'a'.
// Note: You may want initialize the rand seed once in your program.
//
//    rand.Seed(time.Now().UnixNano())
//
func Shuffle(a []string) []string {
	if m := len(a); m > 1 {
		rand.Shuffle(m, func(i, j int) {
			a[i], a[j] = a[j], a[i]
		})
	}
	return a
}

// Trim returns a slice with all the entries of 'a' that don't match string 's'
func Trim(a []string, s string) []string {
	return trimFunc(a, s, streq)
}

// trimFunc runs through 'a' and compares each element with 's' using 'f' comparison
// function (bool). It returns a slice of the elements in 'a' where 'f' returns false.
func trimFunc(a []string, s string, f func(string, string) bool) []string {
	var b []string
	if f == nil {
		return a
	}
	for i := range a {
		if !f(a[i], s) {
			b = append(b, a[i])
		}
	}
	return b
}

// TrimFunc returns a slice with all the entries of 'a' that don't match string 's' that
// satisfy f(s). If func f returns true, the value will be trimmed from 'a'.
func TrimFunc(a []string, f func(string) bool) []string {
	var b []string
	if f == nil {
		return a
	}
	for i := range a {
		if !f(a[i]) {
			b = append(b, a[i])
		}
	}
	return b
}

// TrimPrefix returns a slice with all the entries of 'a' that don't have prefix 'prefix'
func TrimPrefix(a []string, prefix string) []string {
	return trimFunc(a, prefix, strings.HasPrefix)
}

// TrimSuffix returns a slice with all the entries of 'a' that don't have suffix 'suffix'
func TrimSuffix(a []string, suffix string) []string {
	return trimFunc(a, suffix, strings.HasSuffix)
}

// Unshift prepends one or more elements to '*a' and returns the number of entries.
// Note that this function will change the array pointed by 'a'
func Unshift(a *[]string, s ...string) int {
	if s != nil {
		*a = append(s, *a...)
	}
	return len(*a)
}

func fanOutFunc(a, b []string, fof func([]string, string) int) int {
	l := len(b)
	rc := make(chan int, l)
	for i := 0; i < l; i++ {
		go func(s string) {
			rc <- fof(a, s)
		}(b[i])
	}
	ret := -1
	for r := range rc {
		if r != -1 && r < ret {
			ret = r
		}
	}
	return ret
}
