package main

import (
	"log"
)

func main() {
	var strI interface{}
	var intI interface{}

	strI = "hello"
	intI = 20
	log.Println(strI)
	log.Println(intI)

	if value, ok := strI.(string); ok {
		log.Println(value)
	}
	if value, ok := intI.(int); ok {
		log.Println(value)
	}

	// log.Println(strI.(int)) //pacic
	log.Println(strI.(string))
	log.Println(intI.(int))

}
