package model

import (
	"ARTIST_PROJECT_NAME/conf"

	"github.com/coldwind/artist/pkg/ilog"
	"github.com/coldwind/artist/pkg/imysql"
	"github.com/coldwind/artist/pkg/iredis"
	"go.uber.org/zap"
)

var mysqlHandles map[string]*imysql.Service
var redisHandles map[string]*iredis.Service

func Run(mysqlConf *conf.MysqlConf, redisConf *conf.RedisConf) {
	initMysql(mysqlConf)
	initRedis(redisConf)
}

func initMysql(mysqlConf *conf.MysqlConf) {
	mysqlHandles = make(map[string]*imysql.Service)

	for _, cfg := range mysqlConf.Hosts {
		mysqlHandles[cfg.Name] = imysql.New(
			imysql.WithConnection(cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DB),
			imysql.WithLimit(cfg.MaxIdle, cfg.MaxOpen),
			imysql.WithPrefix(cfg.Prefix),
			imysql.WithDebug(cfg.Debug),
		)
	}

	for k, f := range mysqlHandles {
		if err := f.Run(); err != nil {
			ilog.Error("mysql start error", zap.String("key", k), zap.Error(err))
			panic(err)
		} else {
			ilog.Info("mysql started", zap.String("key", k))
		}
	}
}

func initRedis(redisConf *conf.RedisConf) {
	redisHandles = make(map[string]*iredis.Service)

	for _, cfg := range redisConf.Hosts {
		redisHandles[cfg.Name] = iredis.New(
			iredis.WithConnection(cfg.Host, cfg.Port),
			iredis.WithAuth(cfg.Auth),
			iredis.WithLimit(cfg.MaxIdle, cfg.MaxActive),
		)
	}

	for k, f := range redisHandles {
		if err := f.Run(); err != nil {
			ilog.Error("redis start error", zap.String("key", k), zap.Error(err))
			panic(err)
		} else {
			ilog.Info("redis started", zap.String("key", k))
		}
	}
}
