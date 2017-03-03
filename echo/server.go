package main

import (
	"io"
	"log"
	"net"
)

func Server() {
	linsten, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("I am linstening")
	for {
		conn, err := linsten.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// log.Println("it is already connected")
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	if _, err := io.Copy(c, c); err != nil {
		log.Println(err)
	}
	log.Println(c.RemoteAddr())
	c.Close()
}

func main() {
	Server()
}
