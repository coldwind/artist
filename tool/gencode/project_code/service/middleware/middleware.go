package middleware

import "github.com/valyala/fasthttp"

func OutputMiddleware(ctx *fasthttp.RequestCtx) {
	SetCORSHeader(ctx)
}

func SetCORSHeader(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowOrigin, "*")
	ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowMethods, "GET,POST,PUT,DELETE,OPTIONS")
	ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowCredentials, "true")
	ctx.Response.Header.Add(fasthttp.HeaderAccessControlAllowHeaders, "Origin,X-Requested-With,Content-Type,Authorization")
}
