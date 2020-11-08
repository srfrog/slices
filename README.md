# Slices
[![PkgGoDev](https://pkg.go.dev/badge/github.com/srfrog/slices)](https://pkg.go.dev/github.com/srfrog/slices)
[![Go Report Card](https://goreportcard.com/badge/github.com/srfrog/slices?svg=1)](https://goreportcard.com/report/github.com/srfrog/slices)
[![codecov](https://codecov.io/gh/srfrog/slices/branch/master/graph/badge.svg?token=IDUWTIYYZQ)](https://codecov.io/gh/srfrog/slices)
![Build Status](https://github.com/srfrog/slices/workflows/Go/badge.svg)

*Functions that operate on slices. Similar to functions from `package strings` or `package bytes` that have been adapted to work with slices.*

## Features

- [x] Using a thin layer of idiomatic Go; correctness over performance.
- [x] Provide most basic slice operations: index, trim, filter, map
- [x] Some PHP favorites like: pop, push, shift, unshift, shuffle, etc...
- [x] Non-destructive returns (won't alter original slice), except for explicit tasks.

## Quick Start

Install using "go get":

        go get github.com/srfrog/slices

Then import from your source:

        import "github.com/srfrog/slices"

View [example_test.go][1] for an extended example of basic usage and features.

## Documentation

The full code documentation is located at GoDoc:

[http://godoc.org/github.com/srfrog/slices](http://godoc.org/github.com/srfrog/slices)

## Usage

This is a en example showing basic usage.

```go
package main

import(
   "fmt"

   "github.com/srfrog/slices"
)

func main() {
   ss := []string{"Go", "nuts", "for", "Go"}

   foo := slices.Repeat("Go",3)
   fmt.Println(foo)

   fmt.Println(slices.Count(ss, "Go"))

   fmt.Println(slices.Index(ss, "Go"))
   fmt.Println(slices.LastIndex(ss, "Go"))

   if slices.Contains(ss, "nuts") {
      ss = slices.Replace(ss, []string{"Insanely"})
   }
   fmt.Println(ss)

   str := slices.Shift(&ss)
   fmt.Println(str)
   fmt.Println(ss)

   slices.Unshift(&ss, "Really")
   fmt.Println(ss)

   fmt.Println(slices.ToUpper(ss))
   fmt.Println(slices.ToLower(ss))
   fmt.Println(slices.ToTitle(ss))

   fmt.Println(slices.Trim(ss,"Really"))
   fmt.Println(slices.Filter(ss,"Go"))

   fmt.Println(slices.Diff(ss,foo))
   fmt.Println(slices.Intersect(ss,foo))

   fmt.Println(slices.Rand(ss,2))

   fmt.Println(slices.Reverse(ss))
}
```

[1]: https://github.com/srfrog/slices/blob/master/example_test.go
