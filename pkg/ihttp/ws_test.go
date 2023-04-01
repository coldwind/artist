package ihttp

import (
	"fmt"
	"testing"
	"time"

	"github.com/coldwind/artist/pkg/ilog"
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
)

type WsCallback struct {
}

var cb = &WsCallback{}

func TestWs(t *testing.T) {
	ilog.Start("/tmp/", "ARTIST_PROJECT_NAME.log", true, true)

	handle := New(
		WithAddress("", 8899),
	)

	handle.RegisterWS("/", websocket.TextMessage, cb)

	if err := handle.Run(); err != nil {
		t.Error(err)
	}
}

func (w *WsCallback) OnConnect(ctx *fasthttp.RequestCtx, client *WSClient) error {
	// login
	fmt.Println("OnConnect")
	return nil
}

func (w *WsCallback) OnMessage(client *WSClient, msgType int, msg []byte) {
	fmt.Println("OnMessage")
	time.Sleep(3 * time.Second)
	panic("111")
}

func (w *WsCallback) OnClose(client *WSClient) {
	fmt.Println("OnClose")
}
