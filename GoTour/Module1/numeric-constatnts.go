package main

import "fmt"

const (
	Big = 1 << 100 // binary number that is 1 followed by 100 zeros
	Small = Big >> 99 // Shift it right 99 places, so we end upn with 1<<1, or 2
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}