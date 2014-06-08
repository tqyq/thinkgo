package main

import (
	. "github.com/astaxie/beego"
)

func (this *Action) ApiUpload() {
	file, fh, err := this.GetFile("file1")
	Info(file, fh.Filename)
	err = this.SaveToFile("file1", "upload/"+fh.Filename)
	if err != nil {
		Error(err)
		this.EchoJsonErr(err)
	} else {
		this.EchoJsonOk()
	}
}
