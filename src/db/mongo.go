package db

import (
	//	"fmt"
	//	"io"
	//	//	"kodec"
	//	"client"
	//	"net"
	//	"os"
	//	"reflect"
	"labix.org/v2/mgo"
	. "github.com/astaxie/beego"
)

const (
	USER string = "user"
	MSG  string = "msg"
)

var (
	session      *mgo.Session
	databaseName = "cms_go"
)

func Session() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.Dial("localhost")
		if err != nil {
			panic(err) // no, not really
		}
	}
	return session.Clone()
}

func M(collection string, f func(*mgo.Collection)) {
	session := Session()
	defer func() {
	if err := recover(); err != nil {
			Error(err)
		}
		session.Close()
	}()
	c := session.DB(databaseName).C(collection)
	f(c)
}
