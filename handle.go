package goIm

import (
	"fmt"
	"github.com/goinggo/mapstructure"
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

	switch pmd {
	case PMD_SINGLE_SEND_MESSAGE:
		var m = new(SendMessageApi)
		mapDecode(data, m)
		singleLogic.SendMessage(conn, m)
		break
	case PMD_ROOM_CREATE:
		var m = new(CreateRoomApi)
		mapDecode(data, m)
		roomLogic.Create(conn, m)
		break
	case PMD_ROOM_JOIN:
		var m = new(JoinRoomApi)
		mapDecode(data, m)
		roomLogic.Join(conn, m)
		break
	case PMD_ROOM_QUIT:
		var m = new(QuitRoomApi)
		mapDecode(data, m)
		roomLogic.Quit(conn, m)
		break
	case PMD_SINGLE_RECEIPT:
		var m = new(SendReceiptApi)
		mapDecode(data, m)
		singleLogic.SendReceipt(conn, m)
		break
	default:
		break
	}
}

func mapDecode(a interface{}, b interface{}) {
	if err := mapstructure.Decode(a, b); err != nil {
		fmt.Println(err)
	}
}
