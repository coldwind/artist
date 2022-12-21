package iredis

import (
	"testing"

	red "github.com/gomodule/redigo/redis"
)

func TestConn(t *testing.T) {
	var (
		rdb  red.Conn
		err  error
		auth = ""
	)
	pool := &red.Pool{
		Dial: func() (conn red.Conn, e error) {
			rdb, err = red.Dial("tcp", "127.0.0.1:6379")
			t.Log("127.0.0.1:6379")
			if err != nil {
				t.Log("Redis Pool Init failure:", err)
			}

			if auth != "" && err == nil {
				rdb.Do("AUTH", auth)
			}

			return rdb, err
		},
		MaxIdle:     10,
		MaxActive:   10,
		IdleTimeout: 20,
	}
	dial, _ := pool.Dial()
	t.Log(err, "|", dial)

	select {}
}
