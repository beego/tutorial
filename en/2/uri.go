package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {

	// Fixed Router
	beego.Get("/fixed", func(ctx *context.Context) {
		json := "Fixed Router"

		ctx.Output.Json(json, true, true)
	})

	// Default matching  //matching /para1/123    :id = 123  also matching /para1/
	beego.Get("/para1/:id", func(ctx *context.Context) {
		para := ctx.Input.Param(":id")

		json := map[string]string{
			"para": para,
		}

		ctx.Output.Json(json, true, true)
	})

	// Default matching  //matching /para2/123    :id = 123  doesn't match /para2/
	beego.Get("/para2/:id!", func(ctx *context.Context) {
		para := ctx.Input.Param(":id")

		json := map[string]string{
			"para": para,
		}

		ctx.Output.Json(json, true, true)
	})

	// Full matching //matching /para3/path/to/123.html :all= path/to/123.html
	beego.Get("/para3-1/:all", func(ctx *context.Context) {
		para := ctx.Input.Param(":all")

		json := map[string]string{
			"para": para,
		}

		ctx.Output.Json(json, true, true)
	})

	// full matching //matching /para3-2/path/to/123.html :splat=path/to/123.html
	beego.Get("/para3-2/*", func(ctx *context.Context) {
		para := ctx.Input.Param(":splat")

		json := map[string]string{
			"para": para,
		}

		ctx.Output.Json(json, true, true)
	})

	// Regex matching //matching /regex1-1/123
	beego.Get("/regex1-1/:id([0-9]+)", func(ctx *context.Context) {
		para := ctx.Input.Param(":id")

		json := map[string]string{
			"para": para,
		}

		ctx.Output.Json(json, true, true)
	})

	// int type, matching :id is int type, same as regex ([0-9]+)
	beego.Get("/regex1-2/:id:int", func(ctx *context.Context) {
		para := ctx.Input.Param(":id")
		json := map[string]string{
			"para": para,
		}

		ctx.Output.Json(json, true, true)
	})

	// Regex string matching //matching /regex2-1/astaxie :username = astaxie
	beego.Get(`/regex2-1/:username([\w]+)`, func(ctx *context.Context) {
		para := ctx.Input.Param(":username")

		json := map[string]string{
			"para": para,
		}

		ctx.Output.Json(json, true, true)
	})

	// string type, matching :hi is string type。same as ([\w]+)
	beego.Get("/regex2-2/:username:string", func(ctx *context.Context) {
		para := ctx.Input.Param(":username")

		json := map[string]string{
			"para": para,
		}

		ctx.Output.Json(json, true, true)
	})

	// matching //matching /regex3/file/api.xml :path= file/api :ext=xml
	// ** The issue I was facing in the video has been fixed in `beego` develop branch **
	// ** https://github.com/astaxie/beego/commit/3f7e91e6a40edc57596d4d6aa18fb7be1e0cbabb **
	beego.Get("/regex3/*.*", func(ctx *context.Context) {
		path := ctx.Input.Param(":path")
		ext := ctx.Input.Param(":ext")

		json := map[string]string{
			"path": path,
			"ext":  ext,
		}

		ctx.Output.Json(json, true, true)
	})

	// custom regex with prefix //matching :id is regex type, matching pre_123.html, :id = 123
	beego.Get("/pre_:id([0-9]+).html", func(ctx *context.Context) {
		para := ctx.Input.Param(":id")

		json := map[string]string{
			"para": para,
		}

		ctx.Output.Json(json, true, true)
	})

	beego.Run()
}
