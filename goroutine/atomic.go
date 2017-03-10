package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var shutdown int64

func main() {

	wg.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)

	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()

}

func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		// Do we need to shutdown.
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
