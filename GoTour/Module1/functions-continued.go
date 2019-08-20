package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("The addition of 12 and 13 is", add(12, 13))
}
