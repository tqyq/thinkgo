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
	this.CacheTest()
}

func (this *Action) CookieTest() {
	this.Cookie("test", "testcookie")
	this.Data["cookie"] = this.Cookie("test")
}

func (this *Action) CacheTest() {
	Info(this.S("cache"))
	this.S("cache", "testcache", 1)
	Info(this.S("cache"))
	this.Data["cache"] = this.S("cache")
}
