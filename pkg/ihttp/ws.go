package ihttp

import (
	"runtime/debug"
	"sync"
	"sync/atomic"

	"github.com/coldwind/artist/pkg/ilog"
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type WSCallback interface {
	OnConnect(ctx *fasthttp.RequestCtx, client *WSClient) error
	OnMessage(cli *WSClient, msgType int, msg []byte)
	OnClose(client *WSClient)
}

type WS struct {
	sync.RWMutex
	cb      WSCallback
	incrId  int64
	msgType int
	pool    map[int64]*WSClient
}

var (
	MsgTypeText = websocket.TextMessage
	MsgTypeBin  = websocket.BinaryMessage
)

var (
	upgrader = websocket.FastHTTPUpgrader{
		CheckOrigin: func(ctx *fasthttp.RequestCtx) bool {
			return true
		},
	}
)

func (w *WS) wsHandle(ctx *fasthttp.RequestCtx) {
	defer func() {
		if e := recover(); e != nil {
			ilog.Error("[panic]", zap.String("stack", string(debug.Stack())))
		}
	}()

	var cli *WSClient = nil
	upgrader.Upgrade(ctx, func(c *websocket.Conn) {
		cli = &WSClient{
			ConnId:   atomic.AddInt64(&w.incrId, 1),
			conn:     c,
			sendChan: make(chan []byte, 128),
			msgType:  w.msgType,
		}
		err := w.cb.OnConnect(ctx, cli)
		if err == nil {
			// start loop write
			go cli.LoopWrite()

			// add to pool
			w.Lock()
			w.pool[cli.ConnId] = cli
			w.Unlock()

			// read message
			for {
				msgType, message, err := c.ReadMessage()
				if err != nil {
					break
				}
				w.cb.OnMessage(cli, msgType, message)
			}
		}

		if cli != nil {
			// clear client data
			cli.Lock()
			defer cli.Unlock()
			cli.isClose = true
			w.cb.OnClose(cli)
			close(cli.sendChan)

			// delete client from pool
			w.Lock()
			delete(w.pool, cli.ConnId)
			w.Unlock()
		}
	})
}
