package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := x
	temp := 0.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
		if temp == z || math.Abs(temp-z) < 0.000001 {
			break
		}
		fmt.Printf("count: %d temp: %v z: %v diff: %v", i, temp, z, math.Abs(temp-z))
		fmt.Println()
		temp = z
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
