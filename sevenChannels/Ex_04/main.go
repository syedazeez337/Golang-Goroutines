package main

import (
	"fmt"
)

func main() {
	// nil state
	var channel chan int

	fmt.Printf("%v\n", channel)

	channel2 := make(chan string)

	fmt.Printf("%v\n", channel2)

	go func() {
		// send a message
		channel2 <- "Go Message"
	}()

	// receiving a message
	msg := <-channel2
	fmt.Printf("%v\n", msg)

}