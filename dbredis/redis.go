package dbredis

import (
	"bytes"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/ugorji/go/codec"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
)

type RedisHandler interface {
	DoSet(commandName string, key string, val interface{})
	DoSetArgs(commandName string, key string, args ...interface{})
	DoExpire(key string, expireTime uint)
	DoSismember(key string, val string) uint64
	DoGet(commandName string, key string) string
	DoGetStringMap(commandName string, key string) map[string]string
	DoGetStrings(commandName string, key string) []string
	NewScriptSet(fileName string, keyCount int, keyNames []string, args ...interface{})
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
	_prefix     = "im"
	_roomLuaMap = make(map[string]string)
	_luaPrefix  = getPathDir("/lua")
)

/*
	创建连接
*/
func CreateConn() (RedisHandler, error) {
	if conn, err := redis.Dial("tcp", "112.74.61.35:6379"); err == nil {
		conn.Do("AUTH", "112233")
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

/*
	设置过期时间
*/
func (rc *redisConn) DoExpire(key string, expireTime uint) {
	if _, err := rc.conn.Do("EXPIRE", getKeyName(key), expireTime); err != nil {
		fmt.Println(err)
	}
}

/*
	用于一些常规的插入操作
*/
func (rc *redisConn) DoSet(commandName string, key string, val interface{}) {
	if _, err := rc.conn.Do(commandName, getKeyName(key), val); err != nil {
		fmt.Println("DoSet")
	}
}

/*
	用于一些常规的插入操作
*/
func (rc *redisConn) DoSetArgs(commandName string, key string, args ...interface{}) {
	arr := append([]interface{}{getKeyName(key)}, args...)
	if _, err := rc.conn.Do(commandName, arr...); err != nil {
		fmt.Println("DoSetArgs")
	}
}

func (rc *redisConn) DoGet(commandName string, key string) string {
	if result, err := redis.String(rc.conn.Do(commandName, getKeyName(key))); err == nil {
		return result
	} else {
		fmt.Println(err)
		return ""
	}
}

func (rc *redisConn) DoGetStringMap(commandName string, key string) map[string]string {
	if result, err := redis.StringMap(rc.conn.Do(commandName, getKeyName(key))); err == nil {
		return result
	} else {
		fmt.Println(err)
		return make(map[string]string)
	}
}

func (rc *redisConn) DoGetStrings(commandName string, key string) []string {
	if result, err := redis.Strings(rc.conn.Do(commandName, getKeyName(key))); err == nil {
		fmt.Println(result)
		return result
	} else {
		fmt.Println(err)
		return nil
	}
}

/*
	判断是否存在
*/
func (rc *redisConn) DoSismember(key string, val string) uint64 {
	if result, err := redis.Uint64(rc.conn.Do("SISMEMBER", getKeyName(key), val)); err == nil {
		return result
	} else {
		fmt.Println("DoSismember")
		return 0
	}
}

/*
	运行lua脚本
*/
func (rc *redisConn) NewScriptSet(fileName string, keyCount int, keyNames []string, args ...interface{}) {
	if _, ok := _roomLuaMap[fileName]; !ok {
		str, _ := readFile(strings.Join([]string{_luaPrefix, "/", fileName, ".lua"}, ""))
		_roomLuaMap[fileName] = str
	}
	text := _roomLuaMap[fileName]
	var script = redis.NewScript(keyCount, text)
	arr := []interface{}{}
	for _, val := range keyNames {
		arr = append(arr, getKeyName(val))
	}
	for _, val := range args {
		if _, ok := val.([]string); ok {
			arr = append(arr, msgpack(val))
		} else {
			arr = append(arr, val)
		}
	}
	_, err := script.Do(rc.conn, arr...)
	if err != nil {
		fmt.Println(err)
	}
}

func (rc *redisConn) NewScriptGet(keyCount int, src string, keys ...interface{}) {
	var script = redis.NewScript(keyCount, src)
	reply, _ := script.Do(rc.conn, keys...)

	array := reply.([]interface{})
	for i, item := range array {
		if item != nil {
			// 假设 key1，key2 均返回 bulk string
			fmt.Println(i, string(item.([]byte)))
		}
	}
}

/*
	读取文件
*/
func readFile(file string) (string, error) {
	f, _ := os.Open(file)
	bt, _ := ioutil.ReadAll(f)
	return string(bt), nil
}

/*
	获取当前目录绝对地址
*/
func getPathDir(directory string) string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), directory)
}

/*
	Msgpack加码
*/
func msgpack(data interface{}) *bytes.Buffer {
	mh := &codec.MsgpackHandle{}
	buf := &bytes.Buffer{}
	enc := codec.NewEncoder(buf, mh)
	enc.Encode(data)
	return buf
}

/*
	给redis key加前缀
*/
func getKeyName(key string) string {
	arr := []string{_prefix, ":", key}
	return strings.Join(arr, "")
}
