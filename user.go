package main

import (
	. "github.com/astaxie/beego"
)

func (this *Action) UserList() {
	start, rows := this.PageParam("page", "rows")
	m := this.F2m("page", "rows")
	total := D(User).Find(m).Count()
	ps := D(User).Find(m).Skip(start).Limit(rows).Sort("-name").All()
	this.EchoJson(&P{"total": total, "rows": ps})
}

func (this *Action) UserAdd() {
	m := this.F2m()
	D(User).Add(m)
	this.EchoJsonOk()
}

func (this *Action) UserUpdate() {
	m := this.F2m()
	D(User).Save(m)
	this.EchoJsonOk()
}

func (this *Action) UserDel() {
	ids := this.Is("ids[]")
	Debug(ids)
	for _, v := range ids {
		D(User).RemoveId(v)
	}
	this.EchoJsonOk()
}
