package main

import (
	"fmt"
	"time"
)

func main() {
	// calling a function directly
	printIt("direct")

	// goroutine function
	go printIt("goroutine")

	// anonymous goroutine function
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// using time library to delay
	time.Sleep(1 * time.Second)
}

func printIt(s string) {
	for i:=0; i<3; i++ {
		fmt.Printf("%v -> %v\n", i, s)
	}
}