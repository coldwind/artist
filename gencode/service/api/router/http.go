package router

import (
	"ARTIST_PROJECT_NAME/conf"
	"ARTIST_PROJECT_NAME/service/api/code"
	"ARTIST_PROJECT_NAME/service/api/internal"
	"ARTIST_PROJECT_NAME/service/api/middleware"

	"github.com/coldwind/artist/pkg/ihttp"
	"github.com/coldwind/artist/pkg/ilog"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type HttpRouter struct {
	httpConf *conf.HttpConf
	handle   *ihttp.Service
}

type routerMethod struct {
	Handle func(ctx *fasthttp.RequestCtx)
	Filter bool
}

func New(cfg *conf.HttpConf) *HttpRouter {
	return &HttpRouter{
		httpConf: cfg,
	}
}

// Run 启动函数
func (h *HttpRouter) Run() {
	h.handle = ihttp.New(
		ihttp.WithAddress("", h.httpConf.HttpPort),
		ihttp.WithCertificate(h.httpConf.HttpsCertFile, h.httpConf.HttpsKeyFile),
		ihttp.WithRate(h.httpConf.RateLimitPerSec, h.httpConf.RrateLimitCapacity),
	)

	h.Register()

	if err := h.handle.Run(); err != nil {
		ilog.Error("http server start error", zap.Error(err))
	}
}

func (h *HttpRouter) Register() {
	for path := range getHandleList {
		h.handle.Register(path, ihttp.MethodGet, h.PrepareCall)
		ilog.Info("register", zap.String("path", path), zap.String("method", ihttp.MethodGet))

	}

	for path := range postHandleList {
		h.handle.Register(path, ihttp.MethodPost, h.PrepareCall)
		ilog.Info("register", zap.String("path", path), zap.String("method", ihttp.MethodPost))
	}
}

func (h *HttpRouter) PrepareCall(ctx *fasthttp.RequestCtx) {
	h.options(ctx)
	path := string(ctx.URI().Path())
	method := string(ctx.Request.Header.Method())

	if method == "GET" {
		if _, ok := getHandleList[path]; ok {
			if getHandleList[path].Filter {
				if outputCode := h.Filter(ctx); outputCode != code.Success {
					internal.OutputError(ctx, outputCode)
					return
				}
			}

			getHandleList[path].Handle(ctx)
		}
	} else if method == "POST" {
		if _, ok := postHandleList[path]; ok {
			postHandleList[path].Handle(ctx)
		}
	}
}

func (h *HttpRouter) options(ctx *fasthttp.RequestCtx) {
	// 处理OPTIONS
	middleware.SetCORSHeader(ctx)
	ctx.SetStatusCode(fasthttp.StatusAccepted)
}

func (h *HttpRouter) Close() {
}
