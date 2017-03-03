package short

import (
	"testing"
)

func TestNewRedis(t *testing.T) {
	c, err := NewRedis("192.168.3.76:6379")
	if err == nil {
		t.Error(err)
	}
	defer c.Close()
	c.Do("APPEND", "key", "value")
}
