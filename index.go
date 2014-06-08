package main

import (
	. "github.com/astaxie/beego"
)

func (this *Action) Index() {
	Info("Index")
	this.Data["Content"] = "index ..."
	this.Data["Content2"] = "nav ..."
	this.TplNames = "index/index.html"
	this.CookieTest()
}

func (this *Action) CookieTest() {
	this.Cookie("test", "testcookie")
	this.Data["cookie"] = this.Cookie("test")
}
