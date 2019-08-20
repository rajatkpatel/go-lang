package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string

	// %v	the value in a default format
	// %q	a single-quoted character literal safely escaped with Go Syntax
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

}
