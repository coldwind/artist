package ihttp

import (
	"sync/atomic"

	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
)

type WSCallback interface {
	OnConnect(client *WSClient)
	OnMessage(msgType int, msg []byte)
	OnClose()
}

type WS struct {
	cb      WSCallback
	incrId  int64
	msgType int
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
	var cli *WSClient = nil
	upgrader.Upgrade(ctx, func(c *websocket.Conn) {
		cli = &WSClient{
			ConnId:   atomic.AddInt64(&w.incrId, 1),
			conn:     c,
			sendChan: make(chan []byte, 128),
			msgType:  w.msgType,
		}
		w.cb.OnConnect(cli)
		for {
			msgType, message, err := c.ReadMessage()
			if err != nil {
				break
			}
			w.cb.OnMessage(msgType, message)
		}
	})

	if cli != nil {
		cli.Lock()
		defer cli.Unlock()
		cli.isClose = true
		close(cli.sendChan)
		w.cb.OnClose()
	}
}
