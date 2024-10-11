package main

import (
	"fmt"
)

func main() {
	// create a channel variable
	c := make(chan int)

	go func() {
		sum := 0
		for i:=0; i<=100; i++ {
			sum += i
		}
		c <- sum
	}()

	result := <- c
	fmt.Println(result)

}