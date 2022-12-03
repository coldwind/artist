package iredis

import (
	"wnwdkj_ws/boot"
	"wnwdkj_ws/conf"
	"wnwdkj_ws/pkg/logger"

	red "github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

type RedisClass struct {
	RedisPool *red.Pool
}

var instance *RedisClass = nil

func New() *RedisClass {
	if instance == nil {
		instance = &RedisClass{}
	}

	return instance
}

func (r *RedisClass) Run(args *boot.BootArgs) {
	redisConf := conf.New().GetRedisConf()

	r.RedisPool = &red.Pool{
		Dial: func() (conn red.Conn, e error) {
			rdb, err := red.Dial("tcp", redisConf.Host)
			if err != nil {
				logger.Error("Redis Pool Init failure:", zap.Error(err))
			}

			if redisConf.Auth != "" && err == nil {
				rdb.Do("AUTH", redisConf.Auth)
			}

			return rdb, err
		},
		MaxIdle:     redisConf.MaxIdle,
		MaxActive:   redisConf.MaxActive,
		IdleTimeout: 0,
	}
}

func (r *RedisClass) Exec(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
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

func (r *RedisClass) HMSetByMap(key string, hashValue map[string]interface{}) (interface{}, error) {
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

func (r *RedisClass) HGetAll(key string) (map[string]string, error) {
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

func (r *RedisClass) Close() {
	err := r.RedisPool.Close()
	logger.Info("redis closed", zap.Error(err))
}
