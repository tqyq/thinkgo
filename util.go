package main

import (
	"fmt"
	. "github.com/astaxie/beego"
	. "labix.org/v2/mgo/bson"
	"strconv"
	"reflect"
	"regexp"
	"strings"
)

type Util struct {
	Controller
}

func (this *Util) I(key string) interface{} {
	v := this.GetString(key)
	i, err := strconv.Atoi(v)
	if err != nil {
		return v
	} else {
		return i
	}
}

func (this *Util) F2m() *M {
	r := this.Ctx.Request
	r.ParseForm()
	m := &M{}
	for k, v := range r.Form {
		if len(v) == 1 {
			(*m)[k] = v[0]
		} else {
			(*m)[k] = v
		}
	}
	return m
}

func (this *Util) JsonOk(msg ...interface{}) {
	if msg == nil {
		msg = []interface{}{"ok"}
	}
	this.Data["json"] = M{"success": true, "msg": msg[0]}
	this.ServeJson()
}

func (this *Util) Echo(msg ...interface{}) {
	var out string = ""
	for _, v := range msg {
		out += fmt.Sprintf("%v", v)
	}
	this.Ctx.WriteString(out)
}

func (this *Util) EchoJson(m *M) {
	this.Data["json"] = m
	this.ServeJson()
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
