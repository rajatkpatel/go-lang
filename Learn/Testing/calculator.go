package main

import "fmt"

func addition(num int) (result int) {
	result = num + 1
	return result
}

func subtraction(num int) (result int) {
	result = num - 1
	return result
}

func main() {
	fmt.Println("Hello world")
	fmt.Println(addition(2))
	fmt.Println(subtraction(2))
}
