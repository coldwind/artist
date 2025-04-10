package iutils

import (
	"runtime/debug"
)

func GoStartup(fn func(), recoverCall func(e any, stack []byte), forever bool) {
	go func() {
		defer func() {
			if e := recover(); e != nil {
				recoverCall(e, debug.Stack())
				if forever {
					fn()
				}
			}
		}()

		fn()
	}()
}
