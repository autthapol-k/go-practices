package main

import "fmt"

func TryStackingDefer() {
	fmt.Println("Counting...")

	defer fmt.Println("Done")

	for i := range 10 {
		defer fmt.Println(i)
	}
}
