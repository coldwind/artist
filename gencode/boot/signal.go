package boot

import (
	"os"

	"github.com/coldwind/artist/pkg/ilog"
	"github.com/coldwind/artist/pkg/isignal"
	"go.uber.org/zap"
)

func closeSignalListen() {
	defer func() {
		if err := recover(); err != nil {
			ilog.Error("signal listen error", zap.Any("err", err))
		}
	}()

	signal := isignal.New()
	signal.Register(os.Interrupt, func(signal os.Signal, args interface{}) {
		close()
		os.Exit(0)
	})
	signal.Listen()
}
