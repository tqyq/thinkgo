package main

import (
	//	"fmt"
	//	"io"
	//	"net"
	//	"os"
	//	"reflect"
	. "github.com/astaxie/beego"
	"labix.org/v2/mgo"
)

const (
	USER string = "user"
	MSG  string = "msg"
)

var (
	session      *mgo.Session
	databaseName = "cms_go"
)

//type C struct {
//	mgo.Session
//}
//
//func (this *C) C(coll string) *mgo.Session {
//
//}

func Db(collection string) *mgo.Collection {
	session := Session()
	return session.DB(databaseName).C(collection)
}

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

func Mgo(collection string, f func(*mgo.Collection)) {
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
