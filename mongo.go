package main

import (
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
			Error("Mgo", err)
		}
		session.Close()
	}()
	c := session.DB(databaseName).C(collection)
	f(c)
}

type MongoModel struct {
	Cname string
	F     *P  // find/query condition
	Start int // query start at
	Rows  int // query max rows
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
	m.Start = start
	return m
}

func (m *MongoModel) Limit(rows int) *MongoModel {
	m.Rows = rows
	return m
}

func (m *MongoModel) All() *[]P {
	ps := []P{}
	Mgo(m.Cname, func(c *mgo.Collection) {
		q := m.query(c)
		q.All(&ps)
	})
	return &ps
}

func (m *MongoModel) Count() int {
	total := 0
	Mgo(m.Cname, func(c *mgo.Collection) {
		q := m.query(c)
		total, _ = q.Count()
	})
	return total
}

func (m *MongoModel) query(c *mgo.Collection) *mgo.Query {
	q := c.Find(m.F).Skip(m.Start)
	if m.Rows > 0 {
		q = q.Limit(m.Rows)
	}
	return q
}

func (m *MongoModel) Add(docs ...interface{}) error {
	var err error
	Mgo(m.Cname, func(c *mgo.Collection) {
		if len(docs) == 1 {
			c.Insert(docs[0])
		} else {
			err = c.Insert(docs)
		}
	})
	return err
}

func (m *MongoModel) Save(docs ...interface{}) error {
	var err error
	Mgo(m.Cname, func(c *mgo.Collection) {
		if len(docs) == 1 && m.F == nil {
			doc := docs[0]
			id := Field(doc, "_id")
			err = c.UpdateId(id, doc)
		} else {
			err = c.Update(m.F, docs)
		}
	})
	return err
}

func (m *MongoModel) RemoveId(id string) {
}

func (m *MongoModel) Remove(selector interface{}) {
}

type P map[string]interface{}
