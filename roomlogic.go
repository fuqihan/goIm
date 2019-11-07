package goIm

import (
	"fmt"
	"goIm/utils"
	"net"
	"sync"
)

const (
	_joinTypeCreate = 1 // 创建
)

type Roomer interface {
	Create(conn net.Conn, obj *CreateRoomApi)
	Join(conn net.Conn, obj *JoinRoomApi) // 加入房间
	Quit(conn net.Conn, obj *QuitRoomApi) // 退出房间
	//GetRoomInfo(conn net.Conn)            // 获取房间信息
	//GetUserRoomList(conn net.Conn)        // 获取某一个用户的房间列表
}

type room struct {
	createMu sync.Mutex
}

func (r *room) Create(conn net.Conn, obj *CreateRoomApi) {
	r.createMu.Lock()
	defer r.createMu.Unlock()
	// 入库后改成查数据库
	if DBRedisConn.DoSismember(REDIS_ROOM_NAME_LIST, obj.RoomName) == 0 {
		roomId := RoomUUIDGen.GetUint32()
		now := utils.GetTimeNow()
		// 房间加入set、
		DBRedisConn.DoSet("SADD", REDIS_ROOM_LIST, roomId)
		DBRedisConn.DoSet("SADD", REDIS_ROOM_NAME_LIST, obj.RoomName)
		//初始化房间详情
		ma := &utils.MaptoArr{[]interface{}{}}
		ma.Add("createDate", now)
		ma.Add("mode", ROOM_MODE_DEFAULT)
		ma.Add("createUser", obj.UserId)
		ma.Add("name", obj.RoomName)
		DBRedisConn.DoSetArgs("HMSET",
			fmt.Sprintf(REDIS_ROOM_DETAIL, roomId), ma.Arr...)
		// 初始化权限
		DBRedisConn.DoSet("SADD",
			fmt.Sprintf(REDIS_ROOM_ROLE, roomId, ROOM_ROLE_ADMIN), obj.UserId)
		// 初始化房间的用户数据
		DBRedisConn.DoSet("SADD",
			fmt.Sprintf(REDIS_ROOM_USERS, roomId), obj.UserId)
		ma.Clone()
		ma.Add("inviteUser", "")
		ma.Add("currentViewTime", now) // 当前纤细看到的时间
		DBRedisConn.DoSetArgs("HMSET",
			fmt.Sprintf(REDIS_ROOM_USER_INFO, roomId, obj.UserId), ma.Arr...)
		DBRedisConn.DoSet("SADD",
			fmt.Sprintf(REDIS_USER_ROOMS, obj.UserId), roomId)
		SendConnMessageJson(conn, PMD_ROOM_CREATE, SEND_CODE_SUCCESS, "")
	} else {
		SendConnMessageJson(conn, PMD_ROOM_CREATE, SEND_CODE_ERROR, "已存在该房间")
	}
}

func (r *room) Join(conn net.Conn, obj *JoinRoomApi) {
	if DBRedisConn.DoSismember(REDIS_ROOM_LIST, obj.RoomId) != 0 && len(obj.UserIds) != 0 {

	} else {
		SendConnMessageJson(conn, PMD_ROOM_JOIN, SEND_CODE_ERROR, "不存在此房间或用户为空")
	}
}

func (c *room) Quit(conn net.Conn, obj *QuitRoomApi) {

	SendConnMessageJson(conn, PMD_ROOM_QUIT, SEND_CODE_ERROR, "不存在此房间")
}

func NewRoomLogic() Roomer {
	return &room{}
}
