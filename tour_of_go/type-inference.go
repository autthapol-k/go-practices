package main

import "fmt"

func TryTypeInference() {
	v := 42
	f := 3.142
	g := 0.867 + 0.5i
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)
}
