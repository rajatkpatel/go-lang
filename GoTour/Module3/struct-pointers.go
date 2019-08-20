package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	fmt.Println(p.X)
	(*p).X = 40
	fmt.Println((*p).X)
	p.X = 1e9 // same as (*p).X
	fmt.Println(v)
}
