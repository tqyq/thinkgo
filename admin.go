package main

import (
	. "github.com/astaxie/beego"
	"labix.org/v2/mgo"
	. "labix.org/v2/mgo/bson"
)

type AdminController struct {
	Util
}

func (this *AdminController) Prepare() {
	Info("Prepare")
}

func (this *AdminController) Get() {
	this.TplNames = "admin/index.html"
}

func (this *AdminController) Login() {
	this.TplNames = "admin/login.html"
}

func (this *AdminController) Users() {
	start := this.I("start")
	limit := this.I("limit")
	var count int = 0
	Mgo(USER, func(c *mgo.Collection) {
		count, _ = c.Find(M{"start": start, "limit": limit}).Count()
	})

	var ms = []M{}
	Mgo(USER, func(c *mgo.Collection) {
		c.Find(M{"start": start, "limit": limit}).All(&ms)
	})
	this.Data["json"] = M{"rows": &ms, "results": count}
	this.ServeJson()
}

func (this *AdminController) UserAdd() {
	Info("dbname", AppConfig.String("dbname"))
	m := this.F2m()
	Info("m=", *m)
	Mgo(USER, func(c *mgo.Collection) {
		c.Insert(m)
	})
	this.Data["json"] = M{"success": true, "msg": "ok"}
	this.ServeJson()
}

func (this *AdminController) UserDel() {
}
