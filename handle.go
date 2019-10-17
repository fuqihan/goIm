package goIm

import (
	"fmt"
	"net"
	"sync"
)

type UserMap struct {
	user map[string]net.Conn // key remoteIp
	mu   sync.Mutex
}

type ConnMap struct {
	conn map[*net.Conn]string
	mu   sync.Mutex
}

type ReadStruct struct {
	Pmd  int
	Data string
}

var (
	userMap = &UserMap{user: make(map[string]net.Conn)}
	connMap = &ConnMap{conn: make(map[*net.Conn]string)}
)

/**
读取文件处理
*/
func ConnHandle(conn net.Conn, str string) {
	m := new(ReadStruct)
	// TODO  添加配置支持proto
	ParseJson(str, m)
	// 判读
	if m.Pmd == Pmd_LOGIN && m.Data != "" {
		userMap.mu.Lock()
		connMap.mu.Lock()
		// TODO token解析 支持传入fn
		if oriConn, ok := userMap.user[m.Data]; ok {
			EmitMessage(oriConn, "login change")
			delete(connMap.conn, &oriConn)
			//oriConn.Close()
		}
		userMap.user[m.Data] = conn
		connMap.conn[&conn] = m.Data
		userMap.mu.Unlock()
		connMap.mu.Unlock()
		EmitMessage(conn, "asasassa")
		fmt.Println(connMap.conn)
		return
	}
	if _, ok := connMap.conn[&conn]; ok {
		emitMethod(conn, m.Pmd, m.Data)
	}
}

/**
分发请求
*/
func emitMethod(conn net.Conn, pmd int, data string) {
	switch pmd {
	default:
		break
	}
}
