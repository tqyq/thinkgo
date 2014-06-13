package main

import (
	. "github.com/astaxie/beego"
)

func (this *Action) UserList() {
	start, rows := this.PageParam("page", "rows")
	p := this.F2p().Rm("page", "rows").Like("name")
	total := D(User).Find(p).Count()
	ps := D(User).Find(p).Skip(start).Limit(rows).Sort("-name").All()
	this.EchoJson(&P{"total": total, "rows": ps})
}

func (this *Action) UserAdd() {
	p := this.F2p()
	D(User).Add(p)
	this.EchoJsonOk()
}

func (this *Action) UserUpdate() {
	p := this.F2p()
	D(User).Save(p)
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
