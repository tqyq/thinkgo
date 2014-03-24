package main

import (
	//	"fmt"
	//	"io"
	//	//	"kodec"
	//	"client"
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