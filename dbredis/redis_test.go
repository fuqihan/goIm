package dbredis

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateConn(t *testing.T) {
	if conn, err := CreateConn(); err == nil {
		conn.DoSet("SET", "aa", "aaa")
		//aaa :=  conn.DoSismember("aa", "s")
		t1 := time.Now().Unix()
		conn.DoGet("aa")
		t2 := time.Now().Unix()
		fmt.Println(t1)
		fmt.Println(t2 - t1)
		conn.CloneConn()
	} else {
		t.Errorf("连接失败")
	}
}
