package main

import (
	. "github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

var bm cache.Cache

func init() {
	AppConfigPath = "conf/app.conf"
	SetStaticPath("/js", "public/js")
	SetStaticPath("/img", "public/img")
	SetStaticPath("/favicon.ico", "public/img/favicon.ico")
	SetStaticPath("/css", "public/css")
	SetStaticPath("/html", "tpl")
	SetStaticPath("/admin/js", "public/admin/js")
	SetStaticPath("/admin/img", "public/admin/img")
	SetStaticPath("/admin/css", "public/admin/css")
	SetStaticPath("/admin/assets", "public/admin")
	SetLevel(LevelTrace)
	SetLogFuncCall(true)
	bm, _ = cache.NewCache("memory", `{"interval":60}`)
}

func main() {
	InitDb()
	Router("/*", &Action{}, "*:Get")
	Run()
}
