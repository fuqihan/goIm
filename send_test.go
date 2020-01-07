package goIm

import (
	"fmt"
	"net"
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
	op := NewIMOptions()
	go Bootstrap(op)
	time.Sleep(2e9)
	if conn, err := net.Listen("tcp", "0.0.0.0:3000"); err == nil {
		conn.Close()
		t.Errorf("Bootstrap port %q 不应该连接成功", iPort)
	}
}

/*
	文件写测试
*/
func TestRead(t *testing.T) {
	op := NewIMOptions()
	if conn, err := InitTestDial(iPort, op); err == nil {
		strs := []string{"a", "b"}
		go ReadConnMessage(conn, op, func(conn net.Conn, s string, op *IMOptions) {
			//_, _ := strconv.Atoi(s)
			//if a != ERROR_PARSE_JSON {
			//	t.Errorf("返回错误 %s, 应该返回 %d", s, ERROR_PARSE_JSON)
			//}
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
func InitTestDial(port string, op *IMOptions) (net.Conn, error) {
	address := ":" + port
	if conn, err := net.Dial("tcp", address); err == nil {
		return conn, nil
	} else {
		go Bootstrap(op)
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
