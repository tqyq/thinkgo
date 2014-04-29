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

type MongoModel struct {
	Cname string
}

type MongoDb struct {
}

func (*MongoDb) D(name string) (m *MongoModel) {
	return &MongoModel{Cname:name}
}

func (*MongoModel) Find() (m *MongoModel) {
	return m
}

func (*MongoModel) Count() (int64, error) {
	return 0, nil
}
