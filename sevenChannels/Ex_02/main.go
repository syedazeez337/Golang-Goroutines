package main

import (
	"fmt"
	"sync"
	"runtime"
	// "time"
)

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

	wg.Add(2)

	fmt.Println("Starting Goroutines...")
	go func() {
		defer wg.Done()

		// time.Sleep(2 * time.Second)
		for i:='a'; i<='z'; i++ {
			fmt.Printf("%c ", i)
		}
	}()

	go func() {
		defer wg.Done()

		// time.Sleep(1 * time.Second)
		for i:=1; i<=26; i++ {
			fmt.Printf("%d ", i)
		}
	}()

	fmt.Println("Waiting...")
	// goroutines start printing here
	wg.Wait()
	
	fmt.Println()
	fmt.Println("Programme done")
}