package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	bufchan := make(chan int, 10)

	go productor("one", bufchan)
	go productor("two", bufchan)
	go consumer("one", bufchan)

	time.Sleep(time.Second * 5)
}

func consumer(cname string, channel chan int) {
	var randNum int
	for {
		randNum = <-channel
		fmt.Printf("consumer name is %s,rand number is %v\n", cname, randNum)
	}
}

func productor(pname string, channel chan int) {
	for {
		fmt.Printf("productor name is %s\n", pname)
		channel <- rand.Intn(500000)
	}
}
