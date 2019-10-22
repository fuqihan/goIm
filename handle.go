package goIm

import (
	"fmt"
	"github.com/goinggo/mapstructure"
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
)

/**
读取文件处理
*/
func ConnHandle(conn net.Conn, str string) {
	m := new(ReadSApi)
	// TODO  添加配置支持proto
	ParseJson(str, m)
	// 判读
	if m.Pmd == PMD_LOGIN && m.Token != "" {
		localUser.mu.Lock()
		defer localUser.mu.Unlock()
		localConn.mu.Lock()
		defer localConn.mu.Unlock()
		// TODO token解析 支持传入fn
		userId := m.Token
		if oriConn, ok := localUser.user[userId]; ok {
			SendConnMessage(oriConn, "login change")
			delete(localConn.conn, oriConn)
			//oriConn.Close()
		}
		localUser.user[userId] = conn
		localConn.conn[conn] = userId
		SendConnMessage(conn, "login success")
		fmt.Println(localUser.user)
		return
	}
	forRoute(conn, m.Pmd, m.Data)
	//if _, ok := connMap.conn[conn]; ok {
	//	fmt.Println(connMap.conn[conn])
	//
	//	emitMethod(conn, m.Pmd, m.Data)
	//}
}

func SendUserMessage(userId string, str string) {
	if conn, ok := localUser.user[userId]; ok {
		SendConnMessage(conn, str)
	}
}

/**
分发请求
*/
func forRoute(conn net.Conn, pmd int, data interface{}) {

	switch pmd {
	case PMD_SINGLE_SEND_MESSAGE:
		var m SendMessageApi
		mapDecode(data, m)
		singleLogic.SendMessage(conn, localConn.conn[conn], m)
		break
	case PMD_ROOM_JOIN:
		var m JoinRoomApi
		mapDecode(data, m)
		roomLogic.Join(conn, m)
		break
	default:
		break
	}
}

func mapDecode(a interface{}, b interface{}) {
	if err := mapstructure.Decode(a, &b); err != nil {
		fmt.Println(err)
	}
}
