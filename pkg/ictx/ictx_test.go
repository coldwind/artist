package ictx

import (
	"fmt"
	"testing"
	"time"
)

func TestIctx(t *testing.T) {
	pCtx := New()
	go func(ctx *ICtx) {
		sCtx := New(WithParent(ctx))
		go func(ctx *ICtx) {
			<-ctx.Done()
			fmt.Println("sub1 done")
		}(sCtx)
		go func(ctx *ICtx) {
			<-ctx.Done()
			fmt.Println("sub2 done")
		}(sCtx)
		go func(ctx *ICtx) {
			<-ctx.Done()
			fmt.Println("sub3 done")
		}(sCtx)
		go func(ctx *ICtx) {
			<-ctx.Done()
			fmt.Println("sub4 done")
		}(sCtx)
		go func(ctx *ICtx) {
			<-ctx.Done()
			fmt.Println("sub5 done")
		}(sCtx)

		<-pCtx.Done()
		fmt.Println("pCtx done")
	}(pCtx)
	time.Sleep(1 * time.Second)
	pCtx.Cancel()
	select {}
}
