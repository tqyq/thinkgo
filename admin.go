package main

import (
	. "github.com/astaxie/beego"
)

func (this *Action) BeforeAdmin() {
	Info("Call Before Admin")
	this.Echo("over")
	panic("over")
}

