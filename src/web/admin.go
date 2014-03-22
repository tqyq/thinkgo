package web

import (
	//	"db"
	. "github.com/astaxie/beego"
	//	"labix.org/v2/mgo"
//	. "labix.org/v2/mgo/bson"
)

type AdminController struct {
	Controller
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
