package ictx

import (
	"context"
	"sync"

	"github.com/valyala/fasthttp"
)

var (
	TopNode *ICtx = nil
)

type ICtx struct {
	sync.RWMutex
	parent   *ICtx
	children []*ICtx
	done     chan struct{}
	Ctx      context.Context
	FastCtx  *fasthttp.RequestCtx
}

func New(options ...option) *ICtx {
	ctx := &ICtx{
		done:     make(chan struct{}, 1),
		children: make([]*ICtx, 0, 5),
	}
	for _, f := range options {
		f(ctx)
	}

	return ctx
}

func (i *ICtx) Done() <-chan struct{} {
	return i.done
}

func (i *ICtx) Cancel() {
	// 遍历子节点 发送done命令
	for _, sub := range i.children {
		sub.Cancel()
	}
	temp := make([]*ICtx, 0, 5)
	i.children = temp
	close(i.done)
}
