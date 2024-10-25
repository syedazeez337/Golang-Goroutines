package main

import (
	"fmt"
	// "time"
)

func main() {
	// create a channel variable
	c := make(chan int)

	// startTime := time.Now()

	go func() {
		sum := 0
		for i:=0; i<=100000; i++ {
			sum += i
		}
		c <- sum
	}()

	// endTime := time.Now()

	result := <- c
	fmt.Println(result)

	// fmt.Println("Time taken: ", endTime.Sub(startTime))

}