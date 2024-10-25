package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	for i := range 10 {
		wg.Add(1)
		go func1(i)
	}
	go func2()

	// time.Sleep(1 * time.Second)

	wg.Wait()
	fmt.Println("Main function")
}

func func1(n int) {
	defer wg.Done()
	fmt.Println("First function>", n)
}

func func2() {
	defer wg.Done()
	fmt.Println("Second function")
}
