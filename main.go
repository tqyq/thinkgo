package main

import (
	. "github.com/astaxie/beego"
	"reflect"
	"regexp"
	"strings"
)

var controllers []ControllerInterface = []ControllerInterface{&AdminController{}, &UserController{}}

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
	for _, c := range controllers {
		reg, err := regexp.Compile(`.*\.(\w+)Controller`)
		if err != nil {
			Info(err)
		} else {
			match := reg.FindStringSubmatch(reflect.TypeOf(c).String())
			if len(match) > 1 {
				Router("/"+strings.ToLower(match[1])+"/", c)
			}
		}
		AutoRouter(c)
	}
	Run()
}
