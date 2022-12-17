package conf

import (
	"fmt"

	"github.com/coldwind/artist/pkg/icfg"
	"github.com/coldwind/artist/pkg/ilog"

	"go.uber.org/zap"
)

type HttpConf struct {
	Debug              bool   `yaml:"debug"`
	Https              bool   `yaml:"https"`
	HttpsCertFile      string `yaml:"httpsCertFile"`
	HttpsKeyFile       string `yaml:"httpsKeyFile"`
	HttpPort           int    `yaml:"httpPort"`
	RateLimitPerSec    int    `yaml:"rateLimitPerSec"`
	RrateLimitCapacity int    `yaml:"rateLimitCapacity"`
}

func (s *Handle) LoadHttp() {
	path := fmt.Sprintf("%s/%s", s.path, "http.yaml")
	err := icfg.Load(icfg.CfgTypeYaml, "http", path, &HttpConf{})
	if err != nil {
		ilog.Error("get yaml error", zap.String("path", path), zap.Error(err))
	}
}
