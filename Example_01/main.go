package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func main() {
	msg = "Hello world"

	wg.Add(2)
	go updateMessage("Hello, Universe")
	go updateMessage("Hello, Cosmos")
	wg.Wait()

	fmt.Println(msg)
}

func updateMessage(s string) {
	wg.Done()
	msg = s
}