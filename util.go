package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	. "github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/captcha"
	"net/smtp"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Util struct {
	Controller
}

func (this *Util) I(key string) interface{} {
	v := this.GetStrings(key)
	if len(v) == 1 {
		i, err := strconv.Atoi(v[0])
		if err == nil {
			return i
		} else {
			return v[0]
		}
	}
	return v
}

func (this *Util) F2m(exclude ...string) P {
	r := this.Ctx.Request
	r.ParseForm()
	m := P{}
	for k, v := range r.Form {
		if !InArray(exclude, k) {
			if len(v) == 1 {
				if len(v[0]) > 0 {
					m[k] = v[0]
				}
			} else {
				m[k] = v
			}
		}
	}
	return m
}

func (this *Util) EchoJsonOk(msg ...interface{}) {
	if msg == nil {
		msg = []interface{}{"ok"}
	}
	this.Data["json"] = P{"success": true, "msg": msg[0]}
	this.ServeJson()
}

func (this *Util) EchoJsonErr(msg ...interface{}) {
	if msg == nil {
		msg = []interface{}{"err"}
	}
	this.Data["json"] = P{"success": false, "msg": msg[0]}
	this.ServeJson()
}

func (this *Util) Echo(msg ...interface{}) {
	var out string = ""
	for _, v := range msg {
		out += fmt.Sprintf("%v", v)
	}
	this.Ctx.WriteString(out)
}

func (this *Util) EchoJson(m interface{}) {
	this.Data["json"] = m
	this.ServeJson()
}

func (this *Util) PageParam(page string, rows string) (start int, rInt int) {
	rInt = 10
	p := this.I(page)
	r := this.I(rows)
	switch r.(type) {
	case int:
		rInt = this.I(rows).(int)
	}
	switch p.(type) {
	case int:
		start = (this.I(page).(int) - 1) * rInt
	}
	return
}

func (this *Util) Cookie(key string, value ...string) (v string) {
	if len(value) == 0 {
		v = this.Ctx.Input.Cookie(key)
		return
	} else {
		this.Ctx.SetCookie(key, value[0])
		v = value[0]
		return
	}
}

func (this *Util) Redirect(url string) {
	this.Ctx.Redirect(302, url)
}

func (this *Util) Captcha() {
	cpt := captcha.NewWithFilter("/captcha/", bm)
	Debug(cpt)
}

func Md5(s string) (r string) {
	h := md5.New()
	h.Write([]byte(s))
	r = hex.EncodeToString(h.Sum(nil))
	return
}

func S(key string, p ...interface{}) (v interface{}) {
	if len(p) == 0 {
		return bm.Get(key)
	} else {
		if len(p) == 2 {
			var ttl int64
			switch p[1].(type) {
			case int:
				ttl = int64(p[1].(int))
			case int64:
				ttl = p[1].(int64)
			}
			bm.Put(key, p[0], ttl)
		} else if len(p) == 1 {
			bm.Put(key, p[0], 1e9)
		}
		return p[0]
	}
}

func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func Field(i interface{}, fieldName string) string {
	return reflect.ValueOf(i).FieldByName(fieldName).String()
}

func InArray(a []string, e string) bool {
	for _, x := range a {
		if x == e {
			return true
		}
	}
	return false
}

func AutoRoute(controllers ...ControllerInterface) {
	for _, c := range controllers {
		reg, err := regexp.Compile(`.*\.(\w+)Controller`)
		if err != nil {
			Info(err)
		} else {
			match := reg.FindStringSubmatch(reflect.TypeOf(c).String())
			if len(match) > 1 {
				Router("/"+strings.ToLower(match[1])+"/", c)
			}
		}
		AutoRouter(c)
	}
}

type P map[string]interface{}
