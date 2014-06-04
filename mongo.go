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
	User string = "user"
	Msg  string = "msg"
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
	F     *P
}

type MongoDb struct {
}

func (*MongoDb) D(name string) (m *MongoModel) {
	return &MongoModel{Cname: name}
}

func (m *MongoModel) Find(p P) *MongoModel {
	m.F = &p
	return m
}

func (m *MongoModel) Skip(start int) *MongoModel {
	return m
}

func (m *MongoModel) Limit(rows int) *MongoModel {
	return m
}

func (m *MongoModel) All(result interface{}) {
}

func (m *MongoModel) Count() int {
	var total int = 0
	Mgo(m.Cname, func(c *mgo.Collection) {
		total, _ = c.Find(m.F).Count()
	})
	return total
}

func (m *MongoModel) Add(docs ...interface{}) error {
	var err error
	Mgo(m.Cname, func(c *mgo.Collection) {
		err = c.Insert(docs)
	})
	return err
}

func (m *MongoModel) Save(docs ...interface{}) error {
	var err error
	Mgo(m.Cname, func(c *mgo.Collection) {
		err = c.Update(m.F, docs)
	})
	return err
}

func (m *MongoModel) RemoveId(id string) {
}

type P map[string]interface{}
