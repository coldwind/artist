package conf

import (
	"fmt"

	"github.com/coldwind/artist/pkg/icfg"
	"github.com/coldwind/artist/pkg/ilog"
	"go.uber.org/zap"
)

type RedisConf struct {
	Host      string `yaml:"host"`
	Auth      string `yaml:"auth"`
	MaxIdle   int    `yaml:"maxIdle"`
	MaxActive int    `yaml:"maxActive"`
}

func (s *Handle) LoadRedis() {
	path := fmt.Sprintf("%s/%s", s.path, "redis.yaml")
	err := icfg.Load(icfg.CfgTypeYaml, "redis", path, &RedisConf{})
	if err != nil {
		ilog.Error("get yaml error", zap.String("path", path), zap.Error(err))
	}
}
