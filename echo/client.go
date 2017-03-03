package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func Client() {
	conn, err := net.DialTimeout("tcp", "192.168.3.2:8088", 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	buf := make([]byte, 128)
	for i := 0; i < 10000; i++ {
		conn.Write([]byte(fmt.Sprintf("%d", i)))
		conn.Read(buf)
		log.Println(string(buf))
	}

}

func main() {
	Client()
}
