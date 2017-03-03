package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func New(address string) (redis.Conn, error) {
	c, err := redis.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func forSetValue(c redis.Conn) {
	for i := 0; i < 100000; i++ {
		c.Do("SET", fmt.Sprintf("key%d", i), i)
	}
}

func getValue(c redis.Conn) []int {
	datas := []int{}

	for i := 0; i < 100000; i++ {
		key := fmt.Sprintf("key%d", i)
		data, err := redis.Int(c.Do("get", key))
		if err != nil {
			panic(err)
		}
		datas = append(datas, data)
	}
	return datas
}
