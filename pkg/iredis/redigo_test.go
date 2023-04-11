package iredis

import (
	"testing"
)

func TestConn(t *testing.T) {
	service := New(WithConnection("192.168.3.71", 6379))
	err := service.Run()
	t.Log("res", err, "--")
}
