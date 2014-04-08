package main

import (
	. "github.com/astaxie/beego"
)

func init() {
	AppConfigPath = "conf/app.conf"
	SetStaticPath("/js", "public/js")
	SetStaticPath("/img", "public/img")
	SetStaticPath("/css", "public/css")
	SetStaticPath("/html", "tpl")
	SetStaticPath("/admin/js", "public/admin/js")
	SetStaticPath("/admin/img", "public/admin/img")
	SetStaticPath("/admin/css", "public/admin/css")
	SetStaticPath("/admin/assets", "public/admin")
}

func main() {
	Router("/*", &Action{}, "*:Get")
	Run()
}
