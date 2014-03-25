package main

import (
	. "github.com/astaxie/beego"
)

type AdminController struct {
	Util
}

func (this *AdminController) init() {
	Info("Init")
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
