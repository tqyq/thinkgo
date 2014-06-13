package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	. "github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/captcha"
	"net/smtp"
	"reflect"
	"strconv"
	"strings"
)

var DbType, DbHost, DbName, DbUser, DbPwd, DbPort = "", "", "", "", "", 0

type Util struct {
	Controller
}

type DbModel interface {
	Explain() (result interface{})
	Find(p P) (m DbModel)
	Field(s ...string) (m DbModel)
	Limit(rows int) (m DbModel)
	Skip(start int) (m DbModel)
	Sort(s string) (m DbModel)
	Add(docs ...interface{}) error
	All() *[]P
	Count() int
	One() (r interface{})
	RemoveId(id string)
	Save(p P) error
}

func InitDb() {
	DbType = AppConfig.String("db::type")
	DbHost = AppConfig.String("db::host")
	DbName = AppConfig.String("db::name")
	DbUser = AppConfig.String("db::user")
	DbPwd = AppConfig.String("db::pwd")
	DbPort, _ = AppConfig.Int("db::port")
}

func (this *Util) I(key string) interface{} {
	v := this.GetString(key)
	i, err := strconv.Atoi(v)
	if err == nil {
		return i
	} else {
		return v
	}
}

func (this *Util) Is(key string) []string {
	return this.GetStrings(key)
}

func (this *Util) F2p() P {
	r := this.Ctx.Request
	r.ParseForm()
	p := P{}
	for k, v := range r.Form {
		if len(v) == 1 {
			if len(v[0]) > 0 {
				p[k] = v[0]
			}
		} else {
			p[k] = v
		}
	}
	return p
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

func D(name string) (m DbModel) {
	if DbType == "mongo" {
		m = MongoModel{Cname: name}
		return
	}
	return nil
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

//func AutoRoute(controllers ...ControllerInterface) {
//	for _, c := range controllers {
//		reg, err := regexp.Compile(`.*\.(\w+)Controller`)
//		if err != nil {
//			Info(err)
//		} else {
//			match := reg.FindStringSubmatch(reflect.TypeOf(c).String())
//			if len(match) > 1 {
//				Router("/"+strings.ToLower(match[1])+"/", c)
//			}
//		}
//		AutoRouter(c)
//	}
//}

type P map[string]interface{}

func (p P) Like(keys ...string) P {
	if DbType == "mongo" {
		for _, k := range keys {
			v := p[k]
			if v != nil {
				p[k] = MgoLike(fmt.Sprintf("%v", v))
			}
		}
	}
	return p
}

func (p P) Rm(exclude ...string) P {
	for _, k := range exclude {
		delete(p, k)
	}
	return p
}
