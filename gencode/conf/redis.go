package conf

import (
	"fmt"

	"github.com/coldwind/artist/pkg/icfg"
)

type RedisConf struct {
	Hosts []*RedisItem `yaml:"hosts"`
}

type RedisItem struct {
	Name      string `yaml:"name"`
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Auth      string `yaml:"auth"`
	MaxIdle   int    `yaml:"maxIdle"`
	MaxActive int    `yaml:"maxActive"`
}

func (s *Handle) LoadRedis() {
	path := fmt.Sprintf("%s/%s", s.path, "redis.yaml")
	err := icfg.Load(icfg.CfgTypeYaml, "redis", path, &RedisConf{})
	if err != nil {
		panic(err)
	}
}
