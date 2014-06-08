package main

import (
//	. "github.com/astaxie/beego"
)

func (this *Action) UserList() {
	start, rows := this.PageParam("page", "rows")
	m := this.F2m("page", "rows")
	total := this.D(User).Find(m).Count()
	ps := this.D(User).Find(m).Skip(start).Limit(rows).All()
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
