package main

import (
	. "github.com/astaxie/beego"
)

var controllers map[string]ControllerInterface = map[string]ControllerInterface{"admin": &AdminController{}}

func init() {
	AppConfigPath = "conf/app.conf"
	SetStaticPath("/js", "public/js")
	SetStaticPath("/img", "public/img")
	SetStaticPath("/css", "public/css")
	SetStaticPath("/admin/js", "public/admin/js")
	SetStaticPath("/admin/img", "public/admin/img")
	SetStaticPath("/admin/css", "public/admin/css")
	SetStaticPath("/admin/assets", "public/admin")
}

func main() {
	Router("/admin", controllers["admin"])
	for _, c := range controllers {
		AutoRouter(c)
	}
	Run()
}
