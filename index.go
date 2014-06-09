package main

import (
	. "github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"time"
)

func (this *Action) Index() {
	Info("Index")
	this.Data["Content"] = "index ..."
	this.Data["Content2"] = "nav ..."
	this.Captcha()
	this.TplNames = "index/index.html"
	this.cookieTest()
	this.cacheTest()
	go this.httpClientTest()
}

func (this *Action) cookieTest() {
	this.Cookie("test", "testcookie")
	this.Data["cookie"] = this.Cookie("test")
}

func (this *Action) cacheTest() {
	Info(S("cache"))
	S("cache", "testcache", 1)
	Info(S("cache"))
	this.Data["cache"] = S("cache")
}

func (this *Action) httpClientTest() {
	req := httplib.Get("http://www.baidu.com/")
	str, err := req.String()
	if err != nil {
		Error(err)
	}
	Info(len(str))
	str, err = httplib.Post("http://www.baidu.com/").SetTimeout(100*time.Second, 30*time.Second).Param("wd", "go").String()
	if err != nil {
		Error(err)
	}
	Info(len(str))
}
