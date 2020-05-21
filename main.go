package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walk func(*tree.Tree, chan int)
	walk = func(t *tree.Tree, ch chan int) {
		if t != nil {
			walk(t.Left, ch)
			ch <- t.Value
			walk(t.Right, ch)
		}
	}
	walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	var c1, c2 = make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for i := range c1 {
		j := <-c2
		if i != j {
			return false
		}
	}
	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)
	fmt.Println(Same(t1, t2))
}
