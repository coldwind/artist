package iredis

import (
	"time"

	"github.com/coldwind/artist/pkg/ilog"

	red "github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

type Service struct {
	RedisPool   *red.Pool
	addr        string
	auth        string
	maxIdle     int
	maxActive   int
	idleTimeout time.Duration
}

type Option func(*Service)

func New(opts ...Option) *Service {
	s := &Service{}

	for _, f := range opts {
		f(s)
	}

	return s
}

func (r *Service) Run() {
	r.RedisPool = &red.Pool{
		Dial: func() (conn red.Conn, e error) {
			rdb, err := red.Dial("tcp", r.addr)
			if err != nil {
				ilog.Error("Redis Pool Init failure:", zap.Error(err))
			}

			if r.auth != "" && err == nil {
				rdb.Do("AUTH", r.auth)
			}

			return rdb, err
		},
		MaxIdle:     r.maxIdle,
		MaxActive:   r.maxActive,
		IdleTimeout: r.idleTimeout,
	}
}

func (r *Service) Exec(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
	conn := r.RedisPool.Get()
	if err := conn.Err(); err != nil {
		return nil, err
	}

	defer conn.Close()

	params := make([]interface{}, 0)
	params = append(params, key)
	if len(args) > 0 {
		for v := range args {
			params = append(params, v)
		}
	}

	return conn.Do(cmd, params...)
}

func (r *Service) HMSetByMap(key string, hashValue map[string]interface{}) (interface{}, error) {
	conn := r.RedisPool.Get()
	if err := conn.Err(); err != nil {
		return nil, err
	}

	params := make([]interface{}, 0)
	params = append(params, key)
	for k, v := range hashValue {
		params = append(params, k, v)
	}

	return conn.Do("hmset", params...)
}

func (r *Service) HGetAll(key string) (map[string]string, error) {
	conn := r.RedisPool.Get()
	if err := conn.Err(); err != nil {
		return nil, err
	}

	reply, err := conn.Do("hgetall", key)
	if err != nil {
		return nil, err
	}

	mapRes, err := red.StringMap(reply, err)
	if err != nil {
		return nil, err
	}

	return mapRes, nil
}

func (r *Service) Close() {
	err := r.RedisPool.Close()
	ilog.Info("redis closed", zap.Error(err))
}
