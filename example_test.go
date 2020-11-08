// Copyright (c) 2019 srfrog - https://srfrog.me
// Use of this source code is governed by the license in the LICENSE file.

package slices_test

import (
	"fmt"
	"strings"

	"github.com/srfrog/slices"
)

// This example shows basic usage of various functions by manipulating
// the slice 'slc'.
func Example_basic() {
	str := `Don't communicate by sharing memory - share memory by communicating`

	// Split string by spaces into a slice.
	slc := strings.Split(str, " ")

	// Count the number of "memory" strings in slc.
	memories := slices.Count(slc, "memory")
	fmt.Println("Memories:", memories)

	// Split slice into two parts.
	parts := slices.Split(slc, "-")
	fmt.Println("Split:", parts, len(parts))

	// Compare second parts slice with original slc.
	diff := slices.Diff(slc, parts[1])
	fmt.Println("Diff:", diff)

	// Chunk the slice
	chunks := slices.Chunk(parts[0], 1)
	fmt.Println("Chunk:", chunks)

	// Merge the parts
	merge := slices.Merge(chunks...)
	fmt.Println("Merge:", merge)

	// OUTPUT:
	// Memories: 2
	// Split: [[Don't communicate by sharing memory] [share memory by communicating]] 2
	// Diff: [Don't communicate sharing -]
	// Chunk: [[Don't] [communicate] [by] [sharing] [memory]]
	// Merge: [Don't communicate by sharing memory]
}
