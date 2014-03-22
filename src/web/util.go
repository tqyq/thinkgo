package web

import (
	"net/http"
	//	. "github.com/astaxie/beego"
	. "labix.org/v2/mgo/bson"
)

func Req2M(r *http.Request) (*M) {
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
