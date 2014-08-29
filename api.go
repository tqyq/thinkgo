package main

import (
	. "github.com/astaxie/beego"
	"os"
)

func (this *Action) BeforeApi() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	this.Ctx.Output.Header("Access-Control-Allow-Headers", "X-Requested-With, Content-Type")
}

func (this *Action) ApiUpload() {
	file, fh, err := this.GetFile("file1")
	Info(file, fh.Filename)
	wd, _ := os.Getwd()
	err = this.SaveToFile("file1", wd+"/upload/"+fh.Filename)
	if err != nil {
		Error(err)
		this.EchoJsonErr(err)
	} else {
		this.EchoJsonOk()
	}
}

func (this *Action) ApiSendmail() {
	user := "tsyadmin@126.com"
	password := "Admin123"
	host := "smtp.126.com:25"
	to := "tsyadmin@126.com"

	subject := "Test send email by golang"

	body := `
	<html>
	<body>
	<h3>
	"Test send email by golang"
	</h3>
	</body>
	</html>
	`
	Debug("send email")
	err := SendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		Debug("send mail error!")
		Debug(err)
	} else {
		Debug("send mail success!")
	}
	this.EchoJsonOk()
}
