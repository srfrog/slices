// Copyright (c) 2025 srfrog - https://srfrog.dev
// Use of this source code is governed by the license in the LICENSE file.

package slices_test

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
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

func ExampleChunk() {
	// Top 3 programming languages
	data := `1,C,16%,2,Java,13%,3,Python,11%`
	slc := strings.Split(data, ",")

	// Split into chunks of 3.
	chunks := slices.Chunk(slc, 3)
	fmt.Println("Chunks:", chunks)

	// Replace C with Go once.
	chunks[0] = slices.Replace(chunks[0], "C", "Go", 1)
	fmt.Println("Replace:", chunks)

	// OUTPUT:
	// Chunks: [[1 C 16%] [2 Java 13%] [3 Python 11%]]
	// Replace: [[1 Go 16%] [2 Java 13%] [3 Python 11%]]
}

func ExampleShift() {
	slc := []string{"Go", "nuts", "for", "Go"}

	slices.Shift(&slc) // returns "Go"
	fmt.Println("Shift:", slc)

	slices.Unshift(&slc, "Really") // returns 4
	fmt.Println("Unshift:", slc)

	// OUTPUT:
	// Shift: [nuts for Go]
	// Unshift: [Really nuts for Go]
}

func ExampleRandFunc() {
	const (
		chars    = `abcdefghijklmnopqrstuvwxz`
		charsLen = len(chars)
	)

	getIntn := func(max int) int {
		n, _ := crand.Int(crand.Reader, big.NewInt(int64(max)))
		return int(n.Int64())
	}

	getRandStr := func() string {
		var sb strings.Builder

		m := getIntn(charsLen)
		for sb.Len() < m {
			sb.WriteByte(chars[getIntn(charsLen)])
		}

		return sb.String()
	}

	// Generate a slice with 10 random strings.
	slc := slices.RepeatFunc(getRandStr, 10)
	fmt.Println("RepeatFunc len():", len(slc))

	// Pick 3 random elements from slc.
	out := slices.RandFunc(slc, 3, getIntn)
	fmt.Println("RandFunc len():", len(out))

	fmt.Println("ContainsAny:", slices.ContainsAny(slc, out))

	// OUTPUT:
	// RepeatFunc len(): 10
	// RandFunc len(): 3
	// ContainsAny: true
}

func ExampleWalk() {
	var n int

	slc := []string{"Go", "go", "GO"}

	// Count the Go's.
	slices.Walk(slc, func(i int, v string) {
		if strings.EqualFold(v, "Go") {
			n++
		}
	})
	fmt.Println("Count:", n)

	// Convert a slice into a map.
	m := make(map[int]string)
	slices.Walk(slc, func(i int, v string) {
		m[i] = v
	})
	fmt.Println("Mapize:", m)

	// OUTPUT:
	// Count: 3
	// Mapize: map[0:Go 1:go 2:GO]
}
