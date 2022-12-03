package isignal

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/coldwind/artist/pkg/ilog"

	"go.uber.org/zap"
)

type Callback func(s os.Signal, args interface{})

type Handle struct {
	sigChan chan os.Signal
	set     map[os.Signal]Callback
}

func (h *Handle) Register(s os.Signal, f Callback) {
	if _, ok := h.set[s]; !ok {
		h.set[s] = f
		signal.Notify(h.sigChan, s)
	}
}

func (h *Handle) Listen() {
	ilog.Info("isignal listen")
	for {
		if s, ok := <-h.sigChan; ok {
			if s == syscall.SIGURG {
				continue
			}
			ilog.Info("catch signal", zap.String("signal", s.String()))
			if _, ok = h.set[s]; ok {
				ilog.Info("call signal function", zap.String("signal", s.String()))
				h.set[s](s, nil)
			}
		} else {
			break
		}
	}
	ilog.Info("signal exit")
}

func New() *Handle {
	return &Handle{
		sigChan: make(chan os.Signal, 1),
		set:     map[os.Signal]Callback{},
	}
}
