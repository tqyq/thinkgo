package main

import (
	. "github.com/astaxie/beego"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var (
	session *mgo.Session
)

func Session() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.Dial(DbHost)
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
	c := session.DB(DbName).C(collection)
	f(c)
}

type MongoModel struct {
	Cname  string
	F      *P     // find/query condition
	Start  int    // query start at
	Rows   int    // query max rows
	S      string // sort
	Select *P     // select field
}

func (m MongoModel) Find(p P) DbModel {
	m.F = &p
	return m
}

func (m MongoModel) Field(s ...string) DbModel {
	if m.Select == nil {
		m.Select = &P{}
	}
	for _, k := range s {
		(*m.Select)[k] = 1
	}
	return m
}

func (m MongoModel) Skip(start int) DbModel {
	m.Start = start
	return m
}

func (m MongoModel) Limit(rows int) DbModel {
	m.Rows = rows
	return m
}

func (m MongoModel) Sort(s string) DbModel {
	m.S = s
	return m
}

func (m MongoModel) Like(k string, v string) DbModel {
	// TODO
	return m
}

func (m MongoModel) All() *[]P {
	ps := []P{}
	Mgo(m.Cname, func(c *mgo.Collection) {
		q := m.query(c)
		q.All(&ps)
	})
	return &ps
}

func (m MongoModel) One() (r interface{}) {
	p := P{}
	Mgo(m.Cname, func(c *mgo.Collection) {
		q := m.query(c)
		q.One(&p)
	})
	return p
}

func (m MongoModel) Count() int {
	total := 0
	Mgo(m.Cname, func(c *mgo.Collection) {
		q := m.query(c)
		total, _ = q.Count()
	})
	return total
}

func (m MongoModel) Add(docs ...interface{}) error {
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

func (m MongoModel) Save(p P) error {
	var err error
	Mgo(m.Cname, func(c *mgo.Collection) {
		id := p["_id"]
		var oid bson.ObjectId
		switch id.(type) {
		case string:
			oid = bson.ObjectIdHex(id.(string))
		case bson.ObjectId:
			oid = id.(bson.ObjectId)
		}
		p["_id"] = oid
		err = c.UpdateId(oid, p)
		if err != nil {
			Error(err)
		}
	})
	return err
}

func (m MongoModel) RemoveId(id string) {
	Mgo(m.Cname, func(c *mgo.Collection) {
		err := c.RemoveId(bson.ObjectIdHex(id))
		if err != nil {
			Error(err)
		}
	})
}

func (m MongoModel) Remove(selector interface{}) {
}

func (m MongoModel) Explain() (result interface{}) {
	p := P{}
	Mgo(m.Cname, func(c *mgo.Collection) {
		q := m.query(c)
		q.Explain(p)
	})
	return p
}

func (m MongoModel) query(c *mgo.Collection) *mgo.Query {
	q := c.Find(m.F).Skip(m.Start)
	if m.Rows > 0 {
		q = q.Limit(m.Rows)
	}
	if len(m.S) > 0 {
		q = q.Sort(m.S)
	}
	if m.Select != nil {
		q = q.Select(m.Select)
	}
	return q
}
