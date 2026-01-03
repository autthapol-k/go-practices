package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	walkHelper(t, ch)
}

func walkHelper(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	walkHelper(t.Left, ch)
	ch <- t.Value
	walkHelper(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		// If one tree has more values than the other, they're different
		if ok1 != ok2 {
			return false
		}

		// If both channels are closed (ok1 and ok2 are both false), trees are the same!
		if !ok1 {
			return true
		}

		// If values differ, trees are different
		if v1 != v2 {
			return false
		}
	}
}

func TryExerciseEquivalentBinaryTrees() {
	ch := make(chan int)
	t1 := tree.New(1)
	go Walk(t1, ch)
	for range 10 {
		fmt.Println(<-ch)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
