package iredis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	handle *redis.Client
	option *redis.Options
}

type Option func(*Service)

func New(opts ...Option) *Service {
	s := &Service{}

	for _, f := range opts {
		f(s)
	}

	return s
}

func (r *Service) Run() error {
	r.handle = redis.NewClient(r.option)

	return r.handle.Get(context.Background(), "go-redis-testkey").Err()
}

func (r *Service) GetConn() *redis.Client {
	return r.handle
}
