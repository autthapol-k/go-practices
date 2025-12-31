package main

import "fmt"

func TryStructFields() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
