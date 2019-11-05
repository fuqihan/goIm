package goIm

import (
	"fmt"
	"net"
	"strconv"
	"testing"
	"time"
)

var (
	iPort = "3000"
)

/*
	端口重复测试
*/
func TestBootstrapRepeat(t *testing.T) {
	go Bootstrap(iPort)
	time.Sleep(1e9)
	if conn, err := net.Listen("tcp", "127.0.0.1:"+iPort); err == nil {
		conn.Close()
		t.Errorf("Bootstrap port %q 不应该连接失败", iPort)
	}
}

/*
	文件写测试
*/
func TestRead(t *testing.T) {
	if conn, err := InitTestDial(iPort); err == nil {
		strs := []string{"a", "b"}
		go ReadConnMessage(conn, func(conn net.Conn, s string) {
			a, _ := strconv.Atoi(s)
			if a != ERROR_PARSE_JSON {
				t.Errorf("返回错误 %s, 应该返回 %d", s, ERROR_PARSE_JSON)
			}
			return
		})
		for _, str := range strs {
			SendConnMessage(conn, []byte(str))
		}
		time.Sleep(1e9)
	}
}

/*
	初始化客户端连接
*/
func InitTestDial(port string) (net.Conn, error) {
	address := ":" + port
	if conn, err := net.Dial("tcp", address); err == nil {
		return conn, nil
	} else {
		go Bootstrap(port)
		time.Sleep(time.Second)
		if conn1, err1 := net.Dial("tcp", address); err1 == nil {
			return conn1, nil
		} else {
			return nil, err1

		}
	}
}

func IniTestLogin(conn net.Conn, userId string) {
	str := fmt.Sprintf(`{"pmd": %d, "token":"%s"}`, PMD_LOGIN, userId)
	SendConnMessage(conn, []byte(str))
}
