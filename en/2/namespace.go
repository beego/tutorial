package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {

	ns := beego.NewNamespace("/v1").
		Cond(func(ctx *context.Context) bool {
		return true
	}).
		Filter("before", auth).
		Get("/fixed", func(ctx *context.Context) {
		ctx.WriteString("/v1/fixed")
	})

	nsv2 := beego.NewNamespace("/v2").
		Namespace(
		beego.NewNamespace("/alice").
			Get("/:id", func(ctx *context.Context) {
			ctx.WriteString("alice v2 id is " + ctx.Input.Param(":id"))
		}),
		beego.NewNamespace("/bob").
			Get("/:id", func(ctx *context.Context) {
			ctx.WriteString("bob v2 id is " + ctx.Input.Param(":id"))
		}),
	)

	beego.AddNamespace(ns, nsv2)
	beego.Run()
}

func auth(ctx *context.Context) {
}
