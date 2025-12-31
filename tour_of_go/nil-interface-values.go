package main

import "fmt"

type I3 interface {
	M()
}

func TryNilInterfaceValues() {
	var i I3
	describe3(i)
	i.M()
}

func describe3(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
