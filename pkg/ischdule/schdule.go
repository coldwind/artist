package ischdule

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	"github.com/coldwind/artist/pkg/ilog"
	"go.uber.org/zap"
)

type schduleItem struct {
	Ctx      context.Context
	Cancel   context.CancelFunc
	Repeat   int
	Interval time.Duration
	Fn       func()
}

var schduleHandle map[string]*schduleItem
var schduleMapKey sync.RWMutex

func init() {
	schduleHandle = make(map[string]*schduleItem)
	schduleMapKey = sync.RWMutex{}
}

func Register(key string, fn func(), interval time.Duration, repeat int) {
	schduleMapKey.Lock()
	defer schduleMapKey.Unlock()
	if _, ok := schduleHandle[key]; ok {
		// 同一个key如果存在则取消前一个计划任务
		schduleHandle[key].Cancel()
	}

	item := &schduleItem{
		Repeat:   repeat,
		Interval: interval,
		Fn:       fn,
	}
	item.Ctx, item.Cancel = context.WithCancel(context.Background())
	schduleHandle[key] = item

	go schduleRun(key)
}

func Deregister(key string) {
	schduleMapKey.Lock()
	defer schduleMapKey.Unlock()
	res := schduleHandle[key]
	res.Cancel()
	delete(schduleHandle, key)
}

func schduleRun(key string) {
	defer func() {
		// 从schdule 中移出key
		schduleMapKey.Lock()
		delete(schduleHandle, key)
		schduleMapKey.Unlock()

		if e := recover(); e != nil {
			ilog.Error("panic prepareCall", zap.Error(fmt.Errorf("%v", e)), zap.String("trace", string(debug.Stack())))
		}
	}()

	schduleMapKey.RLock()
	res := schduleHandle[key]
	schduleMapKey.RUnlock()
	count := 0
LOOP_SCHDULE:
	for {
		if res.Repeat > 0 {
			count++
			if count > res.Repeat {
				break LOOP_SCHDULE
			}
		}
		select {
		case <-time.After(res.Interval):
			res.Fn()
		case <-res.Ctx.Done():
			break LOOP_SCHDULE
		}
	}
}
