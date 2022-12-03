package ihttp

import (
	"fmt"
	"strings"

	"github.com/coldwind/artist/pkg/ilog"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

type Service struct {
	router        *fasthttprouter.Router
	rateLimiter   *rate.Limiter
	https         bool
	httpsCertFile string
	httpsKeyFile  string
	host          string
	port          int
}

type Option func(opt *Service)

type Method string

func New(opts ...Option) *Service {
	r := &Service{
		router: fasthttprouter.New(),
	}
	for _, f := range opts {
		f(r)
	}

	return r
}

func (h *Service) Register(path, method string, f fasthttp.RequestHandler) {
	method = strings.ToLower(method)
	switch method {
	case MethodGet:
		h.router.GET(path, f)
	case MethodPost:
		h.router.POST(path, f)
	case MethodOptions:
		h.router.OPTIONS(path, f)
	}

}

// Run 启动函数
func (h *Service) Run() {

	var addr = ":8889"
	if h.port != 0 {
		addr = fmt.Sprintf("%s:%d", h.host, h.port)
	}
	ilog.Info("start http server", zap.String("addr", addr))

	var err error = nil
	if h.https {
		err = fasthttp.ListenAndServeTLS(addr, h.httpsCertFile, h.httpsKeyFile, h.router.Handler)
	} else {
		err = fasthttp.ListenAndServe(addr, h.router.Handler)
	}

	if err != nil {
		ilog.Error("start http server error", zap.Error(err))
	}
}
