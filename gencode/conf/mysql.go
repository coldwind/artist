package conf

import (
	"fmt"

	"github.com/coldwind/artist/pkg/icfg"
)

type MysqlConf struct {
	Debug    bool   `yaml:"debug"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
	Prefix   string `yaml:"prefix"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}

func (s *Handle) LoadMysql() {
	path := fmt.Sprintf("%s/%s", s.path, "mysql.yaml")
	err := icfg.Load(icfg.CfgTypeYaml, "mysql", path, &MysqlConf{})
	if err != nil {
		panic(err)
	}
}
