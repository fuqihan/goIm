package goIm

import (
	"fmt"
	"goIm/utils"
	"gopkg.in/fatih/set.v0"
	"net"
	"sync"
)

const (
	joinTypeCreate = 1 // 创建
)

type Roomer interface {
	Join(conn net.Conn, obj *JoinRoomApi) // 加入房间
	Quit(conn net.Conn, obj *QuitRoomApi) // 退出房间
	GetRoomInfo(conn net.Conn)            // 获取房间信息
	GetUserRoomList(conn net.Conn)        // 获取某一个用户的房间列表
}

type localRoom struct {
	createDate int64
	mode       uint32
	users      set.Interface
	role       map[uint32]set.Interface
	mu         sync.Mutex
}

type room struct {
	localRoom map[string]*localRoom
	mu        sync.Mutex
}

func (c *room) Join(conn net.Conn, obj *JoinRoomApi) {
	if data, ok := c.localRoom[obj.RoomName]; ok {
		data.mu.Lock()
		defer data.mu.Unlock()

	} else {
		if obj.Type != joinTypeCreate {
			SendConnMessageJson(conn, PMD_ROOM_JOIN, SEND_CODE_ERROR, "join 创建时type参数错误，或已存在该房间")
			return
		}
		if obj.UserId == "" {
			SendConnMessageJson(conn, PMD_ROOM_JOIN, SEND_CODE_ERROR, "join 创建时UserId不能为空")
			return
		}
		c.localRoom[obj.RoomName] = newRoomMap()
		// 创建者为群超管
		c.localRoom[obj.RoomName].role[ROOM_ROLE_ADMIN] = set.New(set.ThreadSafe)
		c.localRoom[obj.RoomName].role[ROOM_ROLE_ADMIN].Add(obj.UserId)
	}
	if obj.UserId != "" {
		c.localRoom[obj.RoomName].users.Add(obj.UserId)
	} else {
		for _, userId := range obj.UserIds {
			c.localRoom[obj.RoomName].users.Add(userId)
		}
	}
	fmt.Println(c.localRoom[obj.RoomName].role)

}

func (c *room) Quit(conn net.Conn, obj *QuitRoomApi) {
	if data, ok := c.localRoom[obj.RoomName]; ok && obj.UserId != "" {
		data.mu.Lock()
		defer data.mu.Unlock()
		data.users.Remove(obj.UserId)
		SendConnMessageJson(conn, PMD_ROOM_QUIT, SEND_CODE_SUCCESS, "")
		return
	}
	SendConnMessageJson(conn, PMD_ROOM_QUIT, SEND_CODE_ERROR, "不存在此房间")
}

func newRoomMap() *localRoom {
	return &localRoom{
		createDate: utils.GetTimeNow(),
		users:      set.New(set.ThreadSafe),
		mode:       ROOM_MODE_DEFAULT,
		role:       make(map[uint32]set.Interface),
	}
}

func NewRoomLogic() Roomer {
	return &room{
		localRoom: make(map[string]*localRoom),
	}
}
