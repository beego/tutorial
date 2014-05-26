package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {

	ns := beego.NewNamespace("/v1").
		Namespace(
		beego.NewNamespace("/bob").
			Get("/:id", func(ctx *context.Context) {
			ctx.WriteString("shop id is " + ctx.Input.Param(":id"))
		}),
		beego.NewNamespace("/oms").
			Get("/:id", func(ctx *context.Context) {
			ctx.WriteString("oms id is " + ctx.Input.Param(":id"))
		}),
	)

	nsv2 := beego.NewNamespace("/v2").
		Namespace(
		beego.NewNamespace("/bob").
			Get("/:id", func(ctx *context.Context) {
			ctx.WriteString("shop v2 id is " + ctx.Input.Param(":id"))
		}),
		beego.NewNamespace("/oms").
			Get("/:id", func(ctx *context.Context) {
			ctx.WriteString("oms v2 id is " + ctx.Input.Param(":id"))
		}),
	)

	beego.AddNamespace(ns, nsv2)
	beego.Run()
}
