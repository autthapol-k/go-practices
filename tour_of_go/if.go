package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func TryIf() {
	sizeInByte, error := fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(sizeInByte, error)
}
