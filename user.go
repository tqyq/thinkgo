package main

import (
	. "github.com/astaxie/beego"
)

func (this *Action) Index() {
	c := this.D(User).Find(P{"name": "action"}).Count()
	Info(c)
	this.Data["Content"] = "index ..."
	this.Data["Content2"] = "nav ..."
	this.TplNames = "index/index.html"
}

func (this *Action) UserList() {
	start, rows := this.PageParam("page", "rows")
	total := this.D(User).Find(nil).Count()
	ps := this.D(User).Find(nil).Skip(start).Limit(rows).All()
	this.EchoJson(&P{"total": total, "rows": ps})
}

func (this *Action) UserAdd() {
	m := this.F2m()
	this.D(User).Add(m)
	this.EchoJsonOk()
}

func (this *Action) UserUpdate() {
	m := this.F2m()
	this.D(User).Save(m)
	this.EchoJsonOk()
}

func (this *Action) UserDel() {
	ids := this.I("ids[]")
	for _, v := range ids.([]string) {
		this.D(User).RemoveId(v)
	}
	this.EchoJsonOk()
}
