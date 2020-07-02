# Slices
[![GoDoc](https://godoc.org/github.com/srfrog/slices?status.svg)](https://godoc.org/github.com/srfrog/slices)
[![Go Report Card](https://goreportcard.com/badge/github.com/srfrog/slices?svg=1)](https://goreportcard.com/report/github.com/srfrog/slices)
[![Coverage Status](https://coveralls.io/repos/github/srfrog/slices/badge.svg?branch=master)](https://coveralls.io/github/srfrog/slices?branch=master)
[![Build Status](https://travis-ci.com/srfrog/slices.svg?branch=master)](https://travis-ci.com/srfrog/slices)

*Functions that operate on slices. Similar to functions from `package strings` or `package bytes` that have been adapted to work with slices.*


## Features

- Using a thin layer of idiomatic Go; correctness over performance.
- Provide most basic slice operations: index, trim, filter, map
- Some PHP favorites like: pop, push, shift, unshift, shuffle, etc...
- Non-destructive returns (won't alter original slice), except for explicit tasks.

## Quick Start

Install using "go get":

        go get github.com/srfrog/slices

Then import from your source:

        import "github.com/srfrog/slices"

View [example_test.go][2] for an extended example of basic usage and features.

## Documentation

The full code documentation is located at GoDoc:

[http://godoc.org/github.com/srfrog/slices](http://godoc.org/github.com/srfrog/slices)

## Usage

This is a en example showing basic usage.

``` go
package main

import(
   "github.com/srfrog/slices"
   "fmt"
)

// This example shows basic usage of various functions by manipulating
// the slice 'ss'.
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

[1]: https://docs.python.org/3.7/library/stdtypes.html#slices
[2]: https://github.com/srfrog/slices/blob/master/example_test.go
