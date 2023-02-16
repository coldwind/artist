package conf

import (
	"fmt"

	"github.com/coldwind/artist/pkg/icfg"
)

type HttpConf struct {
	EnableStdout       bool   `yaml:"enableStdout"`
	EnableDebug        bool   `yaml:"enableDebug"`
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
		panic(err)
	}
}
