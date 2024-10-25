package main

import (
	"fmt"
	// "time"
)

func main() {

	// startTime := time.Now()

	sum := 0
	for i:=1; i<=100000; i++ {
		sum += i
	}

	fmt.Println(sum)

	// endTime := time.Now()
	// fmt.Println("Time taken: ", endTime.Sub(startTime))
}