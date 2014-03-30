package main

import (
//	. "github.com/astaxie/beego"
)

type IndexController struct {
	Util
}

func (this *IndexController) Get() {
	this.Echo("index", this)
}
