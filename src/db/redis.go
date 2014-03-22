package db


import (
//			"fmt"
	//	"client"
	//	"net"
	//	"os"
	//	"reflect"
	"github.com/garyburd/redigo/redis"
	"time"
	. "github.com/astaxie/beego"
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

func Do(cmd string, v ...interface{}) (reply interface{}, err error) {
	c := pool.Get()
	defer c.Close()
	return c.Do(cmd, v...)
}

func DoString(cmd string, v ...interface{}) (reply string, err error) {
	return redis.String(Do(cmd, v...))
}

func DoInt(cmd string, v ...interface{}) (reply int64, err error) {
	return redis.Int64(Do(cmd, v...))
}

func DoBytes(cmd string, v ...interface{}) (reply []byte, err error) {
	return redis.Bytes(Do(cmd, v...))
}

func C() *redis.Conn {
	c := pool.Get()
	defer c.Close()
	return &c
}

func DoTest() {
c := pool.Get()
	defer c.Close()
	Info(c.Do("set", 11, "test1"));
}

