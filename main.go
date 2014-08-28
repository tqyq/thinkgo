package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

var bm cache.Cache

func init() {
	beego.AppConfigPath = "conf/app.conf"
	beego.SetStaticPath("/js", "public/js")
	beego.SetStaticPath("/img", "public/img")
	beego.SetStaticPath("/favicon.ico", "public/img/favicon.ico")
	beego.SetStaticPath("/css", "public/css")
	beego.SetStaticPath("/html", "tpl")
	beego.SetStaticPath("/admin/js", "public/admin/js")
	beego.SetStaticPath("/admin/img", "public/admin/img")
	beego.SetStaticPath("/admin/css", "public/admin/css")
	beego.SetStaticPath("/admin/assets", "public/admin")
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogFuncCall(true)
	bm, _ = cache.NewCache("memory", `{"interval":60}`)
}

func main() {
	InitDb()
	beego.Router("/", &Action{}, "*:Get")
	beego.Router("/*", &Action{}, "*:Get")
	beego.Run()
}
