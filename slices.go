// Copyright (c) 2019 srfrog - https://srfrog.me
// Use of this source code is governed by the license in the LICENSE file.

// Package slices is a collection of functions to operate with string slices.
// Some functions were adapted from the strings package to work with slices, other
// were ported from PHP 'array_*' function equivalents.
package slices

import (
	"math/rand"
	"strings"
)

// streq just compares 2 strings, used as base comparison.
func streq(s1, s2 string) bool { return s1 == s2 }

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

// Compare returns an integer comparing two slices lexicographically.
// The result will be 0 if a==b, or that a has all values of b.
// The result will be -1 if a < b, or a is shorter than b.
// The result will be +1 if a > b, or a is longer than b.
// A nil argument is equivalent to an empty slice.
func Compare(a, b []string) int {
	return compareFunc(a, b, streq)
}

// Compare returns an integer comparing two slices with func f.
func CompareFunc(a, b []string, f func(string, string) bool) int {
	return compareFunc(a, b, f)
}

// Contains returns true if s is in b, false otherwise
func Contains(a []string, s string) bool {
	return Index(a, s) != -1
}

// ContainsAny returns true if any value in b is in b, false otherwise
func ContainsAny(a, b []string) bool {
	return IndexAny(a, b) != -1
}

// ContainsPrefix returns true if any entry in b has prefix, false otherwise
func ContainsPrefix(a []string, prefix string) bool {
	return indexFunc(a, prefix, strings.HasPrefix) != -1
}

// ContainsSuffix returns true if any entry in b has suffix, false otherwise
func ContainsSuffix(a []string, suffix string) bool {
	return indexFunc(a, suffix, strings.HasSuffix) != -1
}

// Count returns the number of occurrences of s in b
func Count(a []string, s string) int {
	var n int

	for i := range a {
		if streq(a[i], s) {
			n++
		}
	}

	return n
}

// Diff returns a slice with all the elements of b that are not found in b
func Diff(a, b []string) []string {
	return diffFunc(a, b, func(a []string, s string) bool { return !Contains(a, s) })
}

// diffFunc compares the elements of b with those of b using f comparison function (bool).
// It returns a slice of the elements in b that are not found in b which f() == true.
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
	return len(a) == len(b) && Compare(a, b) == 0
}

// EqualFold returns a boolean reporting whether a and b
// are the same length and their values are equal under Unicode case-folding.
func EqualFold(a, b []string) bool {
	return len(a) == len(b) && compareFunc(a, b, strings.EqualFold) == 0
}

// Fill is an alias of Repeat.
func Fill(n int, s string) []string {
	return Repeat(s, n)
}

// filterFunc runs through a and compares each element with func f and
// it returns a slice of the elements in a where f returns true.
func filterFunc(a []string, s string, f func(string, string) bool) []string {
	var b []string

	for i := range a {
		if f(a[i], s) {
			b = append(b, a[i])
		}
	}

	return b
}

// Filter returns a slice with all the elements of b that match string s
func Filter(a []string, s string) []string {
	return filterFunc(a, s, streq)
}

// FilterFunc returns a slice with all the elements of b that match string s that
// satisfy f(s). If func f returns true, the value will be filtered from b.
func FilterFunc(a []string, f func(string) bool) []string {
	if f == nil {
		return nil
	}

	return filterFunc(a, "", func(v string, _ string) bool { return f(v) })
}

// FilterPrefix returns a slice with all the elements of b that have prefix.
func FilterPrefix(a []string, prefix string) []string {
	return filterFunc(a, prefix, strings.HasPrefix)
}

// FilterSuffix returns a slice with all the elements of b that have suffix.
func FilterSuffix(a []string, suffix string) []string {
	return filterFunc(a, suffix, strings.HasSuffix)
}

// Chunk will divide a slice into subslices with size elements into a new 2d slice.
// The last chunk may contain less than size elements. If size less than 1, Chunk returns nil.
func Chunk(a []string, size int) [][]string {
	if size < 1 {
		return nil
	}

	aa := make([][]string, 0, (len(a)+size-1)/size)
	for size <= len(a) {
		a, aa = a[size:], append(aa, a[0:size:size])
	}

	if len(a) > 0 {
		aa = append(aa, a)
	}

	return aa
}

// indexFunc runs through b and compares each element with s using f function.
// It returns the index of the first occurrence of s in b, or -1 if not found.
func indexFunc(a []string, s string, f func(string, string) bool) int {
	for i := range a {
		if f(a[i], s) {
			return i
		}
	}

	return -1
}

// Index returns the index of the first instance of s in b, or -1 if not found
func Index(a []string, s string) int {
	return indexFunc(a, s, streq)
}

// IndexAny returns the index of the first instance of b in b, or -1 if not found
func IndexAny(a, b []string) int {
	return fanOutFunc(a, b, Index)
}

// IndexFunc returns the index into a of the first string satifying f(s), or -1 if not found.
func IndexFunc(a []string, f func(string) bool) int {
	return indexFunc(a, "", func(v string, _ string) bool { return f(v) })
}

// Intersect returns a slice with all the elements of b that are found in b.
func Intersect(a, b []string) []string {
	return diffFunc(a, b, Contains)
}

// InsertAt inserts the values in slice a at index idx.
// This func will append the values if idx doesn't fit in the slice or is negative.
func InsertAt(a []string, idx int, values ...string) []string {
	m, n := len(a), len(values)
	if idx == -1 || idx > m {
		idx = m
	}

	if size := m + n; size <= cap(a) {
		b := a[:size]
		copy(b[idx+n:], a[idx:])
		copy(b[idx:], values)

		return b
	}

	b := make([]string, m+n)
	copy(b, a[:idx])
	copy(b[idx:], values)
	copy(b[idx+n:], a[idx:])

	return b
}

// lastIndexFunc runs through b and compares each element with s using f function.
// It returns the index of the last occurrence of s in b, or -1 if not found.
func lastIndexFunc(a []string, s string, f func(string, string) bool) int {
	for i := len(a) - 1; i >= 0; i-- {
		if f(a[i], s) {
			return i
		}
	}

	return -1
}

// LastIndex returns the index of the last instance of s in b, or -1 if not found
func LastIndex(a []string, s string) int {
	return lastIndexFunc(a, s, streq)
}

// LastIndexAny returns the index of the last instance of b in b, or -1 if not found
func LastIndexAny(a, b []string) int {
	return fanOutFunc(a, b, LastIndex)
}

// LastIndexFunc returns the index into a of the last string satifying f(s), or -1 if none do.
func LastIndexFunc(a []string, f func(string) bool) int {
	return lastIndexFunc(a, "", func(v string, _ string) bool { return f(v) })
}

// LastSearch returns the index of the last entry containing the substring s in b,
// or -1 if not found
func LastSearch(a []string, s string) int {
	return lastIndexFunc(a, s, strings.Contains)
}

// Map returns a new slice with the function 'mapping' applied to each element of b
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

// Merge combines zero or many slices together, while preserving the order of elements.
func Merge(aa ...[]string) []string {
	var a []string

	for i := range aa {
		a = append(a, aa[i]...)
	}

	return a
}

// Pop removes the last element in a and returns it, shortening the slice by one.
// If a is empty returns empty string "".
// Note that this function will change the slice pointed by a.
func Pop(a *[]string) string {
	var s string

	if m := len(*a); m > 0 {
		s, *a = (*a)[m-1], (*a)[:m-1]
	}

	return s
}

// Push appends one or more values to a and returns the number of elements.
// Note that this function will change the slice pointed by a.
func Push(a *[]string, values ...string) int {
	if values != nil {
		*a = append(*a, values...)
	}

	return len(*a)
}

// Repeat returns a slice consisting of n copies of s.
func Repeat(s string, n int) []string {
	a := make([]string, n)
	for i := range a {
		a[i] = s
	}

	return a
}

// Replace returns a copy of the slice b with the first n instances of old replaced by new.
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

// ReplaceAll returns a copy of the slice b with all instances of old replaced by new.
func ReplaceAll(a []string, old, new string) []string {
	return Replace(a, old, new, -1)
}

// Rand returns a new slice with n number of random elements of a.
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

// Reverse returns a slice of the reverse index order elements of b.
func Reverse(a []string) []string {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	return a
}

// Search returns the index of the first entry containing the substring s in b,
// or -1 if not found
func Search(a []string, s string) int {
	return indexFunc(a, s, strings.Contains)
}

// Shift shifts the first element of *a and returns it, shortening the slice by one.
// If *a is empty returns empty string "".
// Note that this function will change the slice pointed by a.
func Shift(a *[]string) string {
	var s string

	if m := len(*a); m > 0 {
		s, *a = (*a)[0], (*a)[1:]
	}

	return s
}

// Shuffle returns a slice with randomized order of elements in b.
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

// Slice returns a subslice of the elements from the slice a as specified by the offset and length parameters.
//
// If offset > 0 the subslice will start at that offset in the slice.
// If offset < 0 the subslice will start that far from the end of the slice.
//
// If length > 0 then the subslice will have up to that many elements in it.
// If length == 0 then the subslice will begin from offset up to the end of the slice.
// If length < 0 then the subslice will stop that many elements from the end of the slice.
// If the slice a is shorter than the length, then only the available elements will be present.
//
// If the offset is larger than the size of the slice, an empty slice is returned.
func Slice(a []string, offset, length int) []string {
	m := len(a)
	if length == 0 {
		length = m
	}

	switch {
	case offset > m:
		return nil
	case offset < 0 && (m+offset) < 0:
		offset = 0
	case offset < 0:
		offset = m + offset
	}

	switch {
	case length < 0:
		length = m - offset + length
	case offset+length > m:
		length = m - offset
	}

	if length <= 0 {
		return nil
	}

	return a[offset : offset+length]
}

// Splice removes a portion of the slice a and replace it with the elements of another.
//
// If offset > 0 then the start of the removed portion is at that offset from the beginning of the slice.
// If offset < 0 then the start of the removed portion is at that offset from the end of the slice.
//
// If length > 0 then that many elements will be removed.
// If length == 0 no elements will be removed.
// If length == size removes everything from offset to the end of slice.
// If length < 0 then the end of the removed portion will be that many elements from the end of the slice.
//
// If b == nil then length elements are removed from a at offset.
// If b != nil then the elements are inserted at offset.
//
func Splice(a []string, offset, length int, b ...string) []string {
	m := len(a)
	switch {
	case offset > m:
		return a
	case offset < 0 && (m+offset) < 0:
		offset = 0
	case offset < 0:
		offset = m + offset
	}

	switch {
	case length < 0:
		length = m - offset + length
	case offset+length > m:
		length = m - offset
	}

	if length <= 0 {
		return a
	}

	return append(a[0:offset], append(b, a[offset+length:]...)...)
}

// slice works almost like strings.genSplit() but for slices.
func split(a []string, sep string, n int) [][]string {
	switch {
	case n == 0:
		return nil
	case sep == "":
		return Chunk(a, 1)
	case n < 0:
		n = Count(a, sep) + 1
	}

	aa, i := make([][]string, n+1), 0
	for i < n {
		m := Index(a, sep)
		if m < 0 {
			break
		}
		aa[i] = a[:m]
		a = a[m+1:]
		i++
	}
	aa[i] = a

	return aa[:i+1]
}

// Split divides a slice a into subslices when any element matches the string sep.
//
// If a does not contain sep and sep is not empty, Split returns a
// 2d slice of length 1 whose only element is a.
//
// If sep is empty, Split returns a 2d slice of all elements in a.
// If a is nil and sep is empty, Split returns an empty slice.
//
// Split is akin to SplitN with a count of -1.
func Split(a []string, sep string) [][]string {
	return split(a, sep, -1)
}

// Split divides a slice a into subslices when n elements match the string sep.
//
// The count determines the number of subslices to return:
//   n > 0: at most n subslices; the last element will be the unsplit remainder.
//   n == 0: the result is nil (zero subslices)
//   n < 0: all subslices
//
// For other cases, see the documentation for Split.
func SplitN(a []string, sep string, n int) [][]string {
	return split(a, sep, n)
}

// trimFunc runs through a and compares each element with func f and
// it returns a slice of the elements in a where f returns false.
func trimFunc(a []string, s string, f func(string, string) bool) []string {
	var b []string

	for i := range a {
		if !f(a[i], s) {
			b = append(b, a[i])
		}
	}

	return b
}

// Trim returns a slice with all the elements of a that don't match string s.
func Trim(a []string, s string) []string {
	return trimFunc(a, s, streq)
}

// TrimFunc returns a slice with all the elements of a that don't match string s that
// satisfy f(s). If func f returns true, the value will be trimmed from a.
func TrimFunc(a []string, f func(value string) bool) []string {
	if f == nil {
		return a
	}

	return trimFunc(a, "", func(v string, _ string) bool { return f(v) })
}

// TrimPrefix returns a slice with all the elements of a that don't have prefix.
func TrimPrefix(a []string, prefix string) []string {
	return trimFunc(a, prefix, strings.HasPrefix)
}

// TrimSuffix returns a slice with all the elements of a that don't have suffix.
func TrimSuffix(a []string, suffix string) []string {
	return trimFunc(a, suffix, strings.HasSuffix)
}

// Unique returns a slice with duplicate values removed.
func Unique(a []string) []string {
	seen := make(map[string]struct{})

	b := a[:0]
	for _, v := range a {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			b = append(b, v)
		}
	}

	return b
}

// Unshift prepends one or more elements to *a and returns the number of elements.
// Note that this function will change the slice pointed by a
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
