package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	tch1, tch2 := make(chan int), make(chan int)

	go Walk(t1, tch1)
	go Walk(t2, tch2)

	for i := 0; i < 10; i++ {
		v1 := <-tch1
		v2 := <-tch2
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	/*
		for i:=0; i<10; i++ {
			fmt.Println(<-ch)
		}
	*/
	fmt.Println(Same(tree.New(1), tree.New(2)))

}
