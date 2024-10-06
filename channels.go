package main

import (
	"fmt"
)

func channels() {
	messages := make(chan string)

	go func() {
		messages <- "Ping"
	}()

	msg := <- messages
	fmt.Println(msg)
}