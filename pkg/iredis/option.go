package iredis

import (
	"fmt"
	"time"
)

func WithConnection(host string, port int) Option {
	return func(s *Service) {
		s.addr = fmt.Sprintf("%s:%d", host, port)
	}
}

func WithAuth(auth string) Option {
	return func(s *Service) {
		s.auth = auth
	}
}

func WithLimit(maxIdle, maxActive int, timeout time.Duration) Option {
	return func(s *Service) {
		s.maxIdle = maxIdle
		s.maxActive = maxActive
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Service) {
		s.idleTimeout = timeout
	}
}
