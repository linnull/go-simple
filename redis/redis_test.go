package redis

import (
	"testing"
)

func TestNew(t *testing.T) {
	conn, err := New("127.0.0.1:6379")
	if err != nil {
		t.Error(err)
	}
	conn.Close()
}

func TestSet(t *testing.T) {
	conn, _ := New("127.0.0.1:6379")
	forSetValue(conn)
	conn.Close()
}

func TestGet(t *testing.T) {
	conn, _ := New("127.0.0.1:6379")
	getValue(conn)
	conn.Close()
}
