package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := 1.0
	temp := 0.0

	if x == 0 {
		fmt.Println("Can not calculate the value of x : ", x)
		return 0
	}
	 
	for i := 0; i < 10; i++ {
		z -= (z * z - x) / (2 * z)
		//fmt.Printf("temp: %v z: %v diff: %v", temp, z, math.Abs(temp - z))
		//fmt.Println()
		if temp == z || math.Abs(temp - z) < 0.000001 {
			break
		}
		temp = z
	}
	return z
}

func main() {
	fmt.Println(sqrt(2))
}