package main

import (
	//			"fmt"
	//	"net"
	//	"os"
	//	"reflect"
	. "github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool = &redis.Pool{
	MaxIdle:     3,
	IdleTimeout: 1 * time.Hour,
	Dial: func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", "127.0.0.1:6379")
		if err != nil {
			Error(err)
			return nil, err
		}
		return c, err
	},
	TestOnBorrow: func(c redis.Conn, t time.Time) error {
		_, err := c.Do("PING")
		return err
	},
}

func C() *redis.Conn {
	c := pool.Get()
	defer c.Close()
	return &c
}
