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
	go Bootstrap(iPort)
	time.Sleep(1e9)
	if conn, err := net.Listen("tcp", "127.0.0.1:"+iPort); err == nil {
		conn.Close()
		t.Errorf("Bootstrap port %q 不应该连接失败", iPort)
	}
}

func TestRead(t *testing.T) {
	if conn, err := InitDial(iPort); err == nil {
		str := "asasaasasasdfksldfjlkfjslfkjslfkjslfjksfjsdlf"
		n, _ := conn.Write([]byte(str))
		fmt.Println(n)
		conn.Close()
	}
}

func InitDial(port string) (net.Conn, error) {
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
