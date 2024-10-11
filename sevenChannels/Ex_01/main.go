package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan string) //unbufferred channel

	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- "Goroutine Message"
		fmt.Printf("Child: signal sent\n")
	}()

	msg := <- ch
	fmt.Printf("Parent: received the signal %s\n", msg)
	wg.Wait()

	fmt.Println("Programme Done")
}