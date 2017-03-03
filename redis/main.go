package main

import (
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "192.168.3.76:6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	c.Do("APPEND", "key", "value")
}
