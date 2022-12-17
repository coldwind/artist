package boot

import (
	"ARTIST_PROJECT_NAME/conf"
	"ARTIST_PROJECT_NAME/service/model"
	"ARTIST_PROJECT_NAME/service/router"

	"github.com/coldwind/artist/pkg/ilog"
)

type BootArgs struct {
	EtcPath string
	LogPath string
}

type BootHandler interface {
	Run()
	Close()
}

func Start(etcPath string, logPath string) {
	// start log
	ilog.Start(logPath, "ARTIST_PROJECT_NAME.log", true)

	// start signal
	go closeSignalListen()

	// load conf
	confHandle := conf.New(etcPath)
	confHandle.Run()
	ilog.Info("conf started")
	model.Run(confHandle.GetMysqlConf())
	ilog.Info("model started")
	route := router.New(confHandle.GetHttpConf())
	route.Run()
}

func Stay() {
	select {}
}

// 优雅关闭调用点
func close() {

}
