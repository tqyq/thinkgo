package main

import (
	. "github.com/astaxie/beego"
	"labix.org/v2/mgo"
	. "labix.org/v2/mgo/bson"
)

func (this *Action) Index() {
	this.D(USER).Find().Count()
	this.Data["Content"] = "index ..."
	this.Data["Content2"] = "nav ..."
	this.TplNames = "index/index.html"
}

func (this *Action) UserList() {
	page := this.I("page").(int)
	rows := this.I("rows").(int)
	start := (page -1) * rows
	var total int = 0
	Mgo(USER, func(c *mgo.Collection) {
		total, _ = c.Find(M{}).Skip(start).Limit(rows).Count()
	})
	var ms = []M{}
	Mgo(USER, func(c *mgo.Collection) {
		c.Find(nil).Skip(start).Limit(rows).All(&ms)
	})
	this.EchoJson(&M{"total": total, "rows": &ms})
}

func (this *Action) UserAdd() {
	m := this.F2m()
	Info("m=", *m)
	Mgo(USER, func(c *mgo.Collection) {
		c.Insert(m)
	})
	this.EchoJsonOk()
}

func (this *Action) UserUpdate() {
	this.Echo("UserUpdate")
}

func (this *Action) UserDel() {
	ids := this.GetStrings("ids[]")
	for _, v := range ids {
		Mgo(USER, func(c *mgo.Collection) {
			c.Remove(M{"_id": ObjectIdHex(v)})
		})
	}
	this.EchoJsonOk()
}
