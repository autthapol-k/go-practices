package main

import "fmt"

func describe4(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func TryEmptyInterface() {
	var i interface{}
	describe4(i)

	i = 42
	describe4(i)

	i = "hello"
	describe4(i)
}
