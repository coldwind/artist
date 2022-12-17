package model

import (
	"ARTIST_PROJECT_NAME/conf"

	"github.com/coldwind/artist/pkg/ilog"
	"github.com/coldwind/artist/pkg/imysql"
	"go.uber.org/zap"
)

var mysqlHandles map[string]*imysql.Service

func Run(cfg *conf.MysqlConf) {
	mysqlHandles = map[string]*imysql.Service{
		"game": imysql.New(
			imysql.WithConnection(cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DB),
			imysql.WithLimit(cfg.MaxIdle, cfg.MaxOpen),
			imysql.WithPrefix(cfg.Prefix),
			imysql.WithDebug(cfg.Debug),
		),
	}

	for k, f := range mysqlHandles {
		if err := f.Run(); err != nil {
			ilog.Error("mysql start error", zap.String("key", k), zap.Error(err))
		} else {
			ilog.Info("mysql started", zap.String("key", k))
		}
	}
}
