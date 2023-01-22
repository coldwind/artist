package conf

import "github.com/coldwind/artist/pkg/icfg"

type Handle struct {
	path string
}

var handler *Handle = nil

func New(etcPath string) *Handle {
	if handler == nil {
		handler = &Handle{
			path: etcPath,
		}
	}

	return handler
}

func GetHandle() *Handle {
	return handler
}

func (s *Handle) Run() {
	s.LoadHttp()
	s.LoadMysql()
	s.LoadRedis()
}

func (s *Handle) GetHttpConf() *HttpConf {
	return icfg.Get("http").(*HttpConf)
}

func (s *Handle) GetMysqlConf() *MysqlConf {
	return icfg.Get("mysql").(*MysqlConf)
}

func (s *Handle) GetRedisConf() *RedisConf {
	return icfg.Get("redis").(*RedisConf)
}

func (s *Handle) Close() {
}
