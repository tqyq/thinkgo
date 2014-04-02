package main

import (
	. "github.com/astaxie/beego"
	"labix.org/v2/mgo"
	. "labix.org/v2/mgo/bson"
)

func (this *Action) UserList() {
	start := this.I("start").(int)
	limit := this.I("limit").(int)
	var count int = 0
	Mgo(USER, func(c *mgo.Collection) {
		count, _ = c.Find(M{}).Skip(start).Limit(limit).Count()
	})
	var ms = []M{}
	Mgo(USER, func(c *mgo.Collection) {
		c.Find(M{}).Skip(start).Limit(limit).All(&ms)
	})
	this.Data["json"] = M{"rows": &ms, "results": count}
	this.ServeJson()
}

func (this *Action) UserAdd() {
	m := this.F2m()
	Info("m=", *m)
	Mgo(USER, func(c *mgo.Collection) {
		c.Insert(m)
	})
	this.JsonOk()
}

func (this *Action) UserUpdate() {
	this.Echo("UserUpdate")
}

func (this *Action) UserDel() {
	this.Echo("UserDel")
}