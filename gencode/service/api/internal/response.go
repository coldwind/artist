package internal

import (
	"ARTIST_PROJECT_NAME/service/api/code"
	"ARTIST_PROJECT_NAME/service/api/middleware"

	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

type ResData map[string]interface{}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Output(ctx *fasthttp.RequestCtx, data interface{}) {
	middleware.OutputMiddleware(ctx)

	res := &Response{
		Code: code.Success.Code,
		Data: data,
	}
	resByte, err := jsoniter.Marshal(res)

	if err != nil {
		return
	}

	ctx.Write(resByte)
}

func OutputError(ctx *fasthttp.RequestCtx, code code.OutputCode) {
	middleware.OutputMiddleware(ctx)
	res := &Response{
		Code: code.Code,
		Msg:  code.Msg,
	}
	resByte, err := jsoniter.Marshal(res)

	if err != nil {
		return
	}

	ctx.Write(resByte)
}
