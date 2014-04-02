package main

import (
	"reflect"
	"regexp"
	"strings"
	. "github.com/astaxie/beego"
)

type Action struct {
	Util
}

func (this *Action) Get() {
	uri := this.Ctx.Input.Uri()
	reg, _ := regexp.Compile(`/(\w+)/?(\w*)/*`)
	match := reg.FindStringSubmatch(uri)
	if len(match) > 1 {
		m1 := match[1]
		m2 := match[2]
		method := strings.ToUpper(m1[0:1]) + m1[1:]
		if len(m2) > 0 {
			method += strings.ToUpper(m2[0:1]) + m2[1:]
		} else {
			m2 = "index"
		}
		defer func() {
			if err := recover(); err != nil {
				Error(err)
				this.TplNames = m1 + "/" + m2 + ".html"
			}
		}()
		reflect.ValueOf(this).MethodByName(method).Call([]reflect.Value{})
	}
}
