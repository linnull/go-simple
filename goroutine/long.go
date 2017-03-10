package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(4)

	fmt.Println("Create Goroutines")

	go printPrime("A")
	go printPrime("B")
	go printPrime("C")
	go printPrime("D")

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nStop Goroutines")
}

func printPrime(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
