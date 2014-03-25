package main

import (
	//	"net/http"
	. "github.com/astaxie/beego"
	. "labix.org/v2/mgo/bson"
	"strconv"
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
