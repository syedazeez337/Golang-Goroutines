package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	
	go func() {
		ch <- "A goroutine message"
		ch <- "Another message"
		ch <- "Third message"
	}()

	for {
		msg := <-ch
		fmt.Println(msg)
	}

	/*
	go func() {
		p := <-ch
		fmt.Println(p)
	}()

	ch <- "Message"
	*/

}