package main

import (
	"fmt"
	. "github.com/astaxie/beego"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Util struct {
	Controller
}

func (this *Util) I(key string) interface{} {
	v := this.GetStrings(key)
	if len(v) == 1 {
		i, err := strconv.Atoi(v[0])
		if err == nil {
			return i
		} else {
			return v[0]
		}
	}
	return v
}

func (this *Util) F2m() *P {
	r := this.Ctx.Request
	r.ParseForm()
	m := &P{}
	for k, v := range r.Form {
		if len(v) == 1 {
			if len(v[0]) > 0 {
				(*m)[k] = v[0]
			}
		} else {
			(*m)[k] = v
		}
	}
	return m
}

func (this *Util) EchoJsonOk(msg ...interface{}) {
	if msg == nil {
		msg = []interface{}{"ok"}
	}
	this.Data["json"] = P{"success": true, "msg": msg[0]}
	this.ServeJson()
}

func (this *Util) Echo(msg ...interface{}) {
	var out string = ""
	for _, v := range msg {
		out += fmt.Sprintf("%v", v)
	}
	this.Ctx.WriteString(out)
}

func (this *Util) EchoJson(m interface{}) {
	this.Data["json"] = m
	this.ServeJson()
}

func (this *Util) PageParam(page string, rows string) (start int, rInt int) {
	rInt = 10
	p := this.I(page)
	r := this.I(rows)
	switch r.(type) {
	case int:
		rInt = this.I(rows).(int)
	}
	switch p.(type) {
	case int:
		start = (this.I(page).(int) - 1) * rInt
	}
	return
}

func Field(i interface{}, fieldName string) string {
	return reflect.ValueOf(i).FieldByName(fieldName).String()
}

func AutoRoute(controllers ...ControllerInterface) {
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
}
