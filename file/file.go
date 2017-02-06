package main

import (
	// "io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("text.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666) //文件没有就创建，文件存在就追加
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString("hello")
}
