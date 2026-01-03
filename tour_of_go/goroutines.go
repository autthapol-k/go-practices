package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func TryGoroutines() {
	go say("world")
	say("hello")

	// say("world")
	// go say("hello")
	// time.Sleep(600 * time.Millisecond)

	// var wg sync.WaitGroup
	// say("world")
	// wg.Add(1) // Tell WaitGroup to wait for 1 goroutine
	// go func() {
	// 	defer wg.Done()
	// 	say("hello")
	// }()
	// wg.Wait() // Wait for all goroutines to complete

	// done := make(chan bool)
	// say("world")
	// go func() {
	// 	say("hello")
	// 	done <- true
	// }()
	// <-done // Wait for signal
}
