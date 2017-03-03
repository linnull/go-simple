package short

import (
	"github.com/garyburd/redigo/redis"
)

func GetRedirectURL(shrot string) string {
	if shrot == "" {
		return ""
	}

	shrot = shrot[1:]

	return "http://www." + shrot + ".com"
}

func NewRedis(address string) (redis.Conn, error) {
	c, err := redis.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return c, nil
}
