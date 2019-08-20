package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for i, _ := range image {
		image[i] = make([]uint8, dx)
	}

	for x, row := range image {
		for y := range row {
			row[y] = uint8((x + y) / 2)
		}
	}

	return image
}

func main() {
	pic.Show(Pic)
}
