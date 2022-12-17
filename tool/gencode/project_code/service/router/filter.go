package router

import (
	"ARTIST_PROJECT_NAME/service/code"

	"github.com/valyala/fasthttp"
)

func (h *HttpRouter) Filter(ctx *fasthttp.RequestCtx) code.OutputCode {
	return code.Success
}
