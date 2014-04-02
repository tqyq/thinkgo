package main

import (
	. "github.com/astaxie/beego"
	"reflect"
	"regexp"
	"strings"
)

type Action struct {
	Util
}

func (this *Action) Get() {
	uri := this.Ctx.Input.Uri()
	uri = strings.ToLower(uri)
	reg, _ := regexp.Compile(`/(\w+)/?(\w*)/*`)
	match := reg.FindStringSubmatch(uri)
	var method, m1, m2 string = "", "Index", "Index"
	if len(match) > 1 {
		m1 = match[1]
		method = strings.ToUpper(m1[0:1]) + m1[1:]
		if len(match[2]) > 0 {
			m2 = match[2]
			method += strings.ToUpper(m2[0:1]) + m2[1:]
		}
	} else {
		method = m1 + m2
	}
	defer func() {
		if err := recover(); err != nil {
			Error(err)
			this.Echo(err)
		}
	}()
	Debug("method ", method)
	v := reflect.ValueOf(this).MethodByName(method)
	if v.IsValid() {
		v.Call([]reflect.Value{})
	} else {
		this.TplNames = strings.ToLower(m1) + "/" + strings.ToLower(m2) + ".html"
	}
	
}
