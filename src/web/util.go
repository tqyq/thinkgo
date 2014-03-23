package web

import (
//	"net/http"
		. "github.com/astaxie/beego"
	. "labix.org/v2/mgo/bson"
)

type Util struct {
Controller
}

func (this *Util) I(key string) string {
	return this.GetString(key)
}

func (this *Util) F2m() (*M) {
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
