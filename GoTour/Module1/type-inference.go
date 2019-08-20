package main

import "fmt"

func main() {
	// change the value of v and see that output will be changed
	// value to change like for int -> 42, float -> 42.1, complex -> 42.1 + 0.2i
	v := 42
	fmt.Printf("V is of type %T\n", v)
}
