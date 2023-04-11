package iredis

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestConn(t *testing.T) {
	service := New(WithConnection("192.168.3.72", 6379))
	err := service.Run()
	t.Log("res", err, "--")
	slice := service.GetConn().ZRangeByScore(context.Background(), "test_zsort", &redis.ZRangeBy{
		Min:    fmt.Sprintf("%d", 2),
		Max:    fmt.Sprintf("%d", 5),
		Offset: int64(0),
		Count:  int64(20),
	})
	t.Log(slice.Result())
}
