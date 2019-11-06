package dbredis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strings"
)

type RedisHandler interface {
	DoSet(commandName string, key string, val string)
	DoSetArgs(commandName string, key string, args ...interface{})
	DoExpire(key string, expireTime uint)
	DoSismember(key string, val string) uint64
	DoGet(key string) string
	CloneConn()
}

type ConnOptions struct {
	host     string
	port     uint32
	password string
}

type redisConn struct {
	conn redis.Conn
}

var (
	_prefix = "im"
)

/*
	创建连接
*/
func CreateConn() (RedisHandler, error) {
	if conn, err := redis.Dial("tcp", "112.74.61.35:6379"); err == nil {
		return &redisConn{conn: conn}, nil
	} else {
		fmt.Println(err)
		return nil, err
	}
}

/*
	关闭连接
*/
func (rc *redisConn) CloneConn() {
	_ = rc.conn.Close()
}

func (rc *redisConn) DoExpire(key string, expireTime uint) {
	if _, err := rc.conn.Do("EXPIRE", getKeyName(key), expireTime); err != nil {
		fmt.Println(err)
	}
}

/*
	用于一些常规的插入操作
*/
func (rc *redisConn) DoSet(commandName string, key string, val string) {
	if _, err := rc.conn.Do(commandName, getKeyName(key), val); err != nil {
		fmt.Println(err)
	}
}

/*
	用于一些常规的插入操作
*/
func (rc *redisConn) DoSetArgs(commandName string, key string, args ...interface{}) {
	arr := append([]interface{}{getKeyName(key)}, args...)
	if _, err := rc.conn.Do(commandName, arr...); err != nil {
		fmt.Println(err)
	}
}

func (rc *redisConn) DoGet(key string) string {
	if result, err := redis.String(rc.conn.Do("GET", getKeyName(key))); err == nil {
		fmt.Println(result)
		return result
	} else {
		fmt.Println(err)
		return ""
	}
}

func (rc *redisConn) DoSismember(key string, val string) uint64 {
	if result, err := redis.Uint64(rc.conn.Do("SISMEMBER", getKeyName(key), val)); err == nil {
		return result
	} else {
		fmt.Println(err)
		return 0
	}
}

func getKeyName(key string) string {
	arr := []string{_prefix, ":", key}
	return strings.Join(arr, "")
}
