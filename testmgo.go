package main

import (
	"fmt"
	. "github.com/astaxie/beego"
	"labix.org/v2/mgo"
	. "labix.org/v2/mgo/bson"
)

type User struct {
	Id     ObjectId `bson:"_id"`
	Name   string
	Passwd string
	Img    []byte
}

func crud() {
	var m = M{}
	cname := "test"
	// insert
	for i := 0; i < 3; i++ {
		u := User{NewObjectId(), cname + fmt.Sprintf("%d", i), "xxx", []byte{byte(i)}}
		Mgo(cname, func(c *mgo.Collection) {
			c.Insert(u)
			Info("insert", u)
			//c.Insert(M{"_id": i, "foo": "bar"})
		})
	}

	// upsert
	Mgo(cname, func(c *mgo.Collection) {
		info, _ := c.Upsert(M{"name": "upsert"}, M{"name": "upsert"})
		Info("upsert", info)
	})
	Mgo(cname, func(c *mgo.Collection) {
		c.Find(M{"name": "upsert"}).One(&m)
		Info("find upsert", m)
	})

	// count
	Mgo(cname, func(c *mgo.Collection) {
		cn, _ := c.Count()
		Info("count", cn)
	})
	Mgo(cname, func(c *mgo.Collection) {
		cn, _ := c.Find(M{"name": "test0"}).Count()
		Info("count by name", cn)
	})

	// find by column
	var users = []User{}
	Mgo(cname, func(c *mgo.Collection) {
		c.Find(M{"name": "test0"}).All(&users)
		Info("users", users)
	})
	id := users[0].Id
	Info("convert", id.String(), id.Hex())

	// sort
	var ms = []M{}
	Mgo(cname, func(c *mgo.Collection) {
		c.Find(M{}).Sort("-name", "_id").All(&ms)
		Info("sort", ms)
	})
	Mgo(cname, func(c *mgo.Collection) {
		c.Find(M{}).Sort("-name").One(m)
		Info("sort one", m)
	})
	Mgo(cname, func(c *mgo.Collection) {
		c.Find(M{}).Sort("-name").Limit(1).All(&ms)
		Info("sort limit", ms)
	})

	// skip
	Mgo(cname, func(c *mgo.Collection) {
		c.Find(M{}).Skip(1).All(&ms)
		Info("skip", ms)
	})

	// find one
	//	objectId := ObjectIdHex(id.(string))
	user := new(User)
	Mgo(cname, func(c *mgo.Collection) {
		c.Find(M{"_id": id}).One(user)
		Info("find one", user, "user img", string(user.Img))
	})
	Mgo(cname, func(c *mgo.Collection) {
		c.Find(M{"_id": id}).One(m)
		Info("find one m", m, m["name"], m["img"])
	})
	Mgo(cname, func(c *mgo.Collection) {
		c.Find(M{}).One(m)
		Info("find one m2", m, m["name"], m["img"])
	})

	// update
	m = M{}
	m["img"] = []byte("gaga")
	Mgo(cname, func(c *mgo.Collection) {
		c.Update(M{"_id": id},
			M{"$set": m})
	})
	Mgo(cname, func(c *mgo.Collection) {
		c.Update(M{"_id": id},
			M{"$set": M{
				"name": "Jimmy",
				"age":  34,
			}})
	})
	Mgo(cname, func(c *mgo.Collection) {
		c.Find(M{"name": "Jimmy"}).One(m)
		Info("updated", m["name"], m["age"], string(m["img"].([]byte)))
	})

	// push & pull
	Mgo(cname, func(c *mgo.Collection) {
		c.Update(M{"_id": id},
			M{"$push": M{
				"interests": "Golang",
			}})
	})

	// remove
	Mgo(cname, func(c *mgo.Collection) {
		c.Remove(M{"name": "Jimmy"})
	})

	user = new(User)
	Mgo(cname, func(c *mgo.Collection) {
		c.Find(M{"_id": id}).One(user)
		Info("not found", user)
	})

	// remove all
	Mgo(cname, func(c *mgo.Collection) {
		c.RemoveAll(M{})
	})
}
