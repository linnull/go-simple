package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	data := []byte("123456")

	sum0 := md5.Sum(data)
	fmt.Println(sum0)

	sum1 := sha1.Sum(data)
	fmt.Println(sum1)

	sum2 := sha256.Sum256(data)
	fmt.Println(sum2)

	sum3 := sha512.Sum512(data)
	fmt.Println(sum3)
}
