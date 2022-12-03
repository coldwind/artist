package ictx

import "time"

type option func(ctx *ICtx)

func WithParent(parent *ICtx) option {
	return func(ctx *ICtx) {
		if ctx != nil {
			ctx.parent = parent
			parent.Lock()
			defer parent.Unlock()
			parent.children = append(parent.children, ctx)
		}
	}
}

func WithTimeout(t time.Duration) option {
	return func(ctx *ICtx) {
		if ctx != nil {
			go func() {
				<-time.After(t)
				ctx.Cancel()
			}()
		}
	}
}
