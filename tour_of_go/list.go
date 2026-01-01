package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func TryList() {
	list := List[int]{
		next: &List[int]{
			next: &List[int]{
				val: 10,
			},
			val: 20,
		},
		val: 20,
	}

	// Print all values in the list
	current := &list
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
}
