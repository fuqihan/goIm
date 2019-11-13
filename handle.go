package goIm

import (
	"fmt"
	"goIm/utils"
	"net"
	"sync"
)

type LocalUser struct {
	user map[string]net.Conn // key remoteIp
	mu   sync.Mutex
}

type LocalConn struct {
	conn map[net.Conn]string
	mu   sync.Mutex
}

type ReadSApi struct {
	Pmd   int
	Token string
	Data  interface{}
}

var (
	localUser   = &LocalUser{user: make(map[string]net.Conn)}
	localConn   = &LocalConn{conn: make(map[net.Conn]string)}
	singleLogic = NewSingleLogic()
	roomLogic   = NewRoomLogic()
	pmdMap      = newPmdMap()
)

/**
读取文件处理
*/
func ConnHandle(conn net.Conn, str string) {
	m := new(ReadSApi)
	// TODO  添加配置支持proto
	if err := utils.ParseJson(str, m); err != nil {
		SendConnMessageJson(conn, 0, 0, "json格式错误")
		return
	}
	// 判读
	if m.Pmd == PMD_LOGIN && m.Token != "" {
		localUser.mu.Lock()
		defer localUser.mu.Unlock()
		localConn.mu.Lock()
		defer localConn.mu.Unlock()
		// TODO token解析 支持传入fn
		userId := m.Token
		if oriConn, ok := localUser.user[userId]; ok {
			SendConnMessageJson(oriConn, PMD_LOGIN, SEND_CODE_SUCCESS, "login change")
			delete(localConn.conn, oriConn)
			//oriConn.Close()
		}
		localUser.user[userId] = conn
		localConn.conn[conn] = userId
		SendConnMessageJson(conn, PMD_LOGIN, SEND_CODE_SUCCESS, "login success")
		fmt.Println(localUser.user)
		return
	}
	if _, ok := localConn.conn[conn]; ok {
		forRoute(conn, m.Pmd, m.Data)
	}
}

func SendUserMessage(userId string, str string) {
	if conn, ok := localUser.user[userId]; ok {
		SendConnMessageStr(conn, str)
	}
}

/**
分发请求
*/
func forRoute(conn net.Conn, pmd int, data interface{}) {
	if fn, ok := pmdMap[pmd]; ok {
		fn(conn, data)
	}
}

func newPmdMap() map[int]func(conn net.Conn, m interface{}) {
	m := make(map[int]func(conn net.Conn, m interface{}))

	m[PMD_SINGLE_SEND_MESSAGE] = singleLogic.SendMessage
	m[PMD_SINGLE_RECEIPT] = singleLogic.SendReceipt
	m[PMD_ROOM_CREATE] = roomLogic.Create
	m[PMD_ROOM_QUIT] = roomLogic.Quit
	m[PMD_ROOM_INFO] = roomLogic.GetRoomInfo

	return m
}
