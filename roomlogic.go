package goIm

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"goIm/utils"
	"net"
	"sync"
)

const (
	_joinTypeCreate = 1 // 创建
)

type Roomer interface {
	Create(conn net.Conn, data interface{})
	Join(conn net.Conn, data interface{})        // 加入房间
	Quit(conn net.Conn, data interface{})        // 退出房间
	GetRoomInfo(conn net.Conn, data interface{}) // 获取房间信息
	//GetUserRoomList(conn net.Conn)        // 获取某一个用户的房间列表
}

type room struct {
	createMu sync.Mutex
}

func (r *room) Create(conn net.Conn, data interface{}) {
	obj := new(CreateRoomApi)
	mapstructure.Decode(data, obj)
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

func (r *room) Join(conn net.Conn, data interface{}) {
	obj := new(JoinRoomApi)
	mapstructure.Decode(data, obj)
	if DBRedisConn.DoSismember(REDIS_ROOM_LIST, obj.RoomId) != 0 && len(obj.UserIds) != 0 {
		now := utils.GetTimeNow()

		keyNames := []string{REDIS_ROOM_USERS, REDIS_ROOM_USER_INFO, REDIS_USER_ROOMS}
		count := len(keyNames)
		DBRedisConn.NewScriptSet("room_join",
			count, keyNames, obj.RoomId, obj.UserIds, now)
		SendConnMessageJson(conn, PMD_ROOM_JOIN, SEND_CODE_SUCCESS, "")
	} else {
		SendConnMessageJson(conn, PMD_ROOM_JOIN, SEND_CODE_ERROR, "不存在此房间或用户为空")
	}
}

func (r *room) Quit(conn net.Conn, data interface{}) {
	obj := new(QuitRoomApi)
	mapstructure.Decode(data, obj)
	if DBRedisConn.DoSismember(REDIS_ROOM_LIST, obj.RoomId) != 0 && len(obj.UserIds) != 0 {

		keyNames := []string{REDIS_ROOM_USERS, REDIS_ROOM_USER_INFO, REDIS_USER_ROOMS}
		count := len(keyNames)
		DBRedisConn.NewScriptSet("room_quit",
			count, keyNames, obj.RoomId, obj.UserIds)
		SendConnMessageJson(conn, PMD_ROOM_QUIT, SEND_CODE_SUCCESS, "")
	} else {
		SendConnMessageJson(conn, PMD_ROOM_QUIT, SEND_CODE_ERROR, "不存在此房间")
	}
}

func (r *room) GetRoomInfo(conn net.Conn, data interface{}) {
	obj := new(GetRoomInfoApi)
	mapstructure.Decode(data, obj)
	if DBRedisConn.DoSismember(REDIS_ROOM_LIST, obj.RoomId) != 0 {
		str := DBRedisConn.DoGet("HGET", fmt.Sprintf(REDIS_ROOM_DETAIL, obj.RoomId))
		SendConnMessageJson(conn, PMD_ROOM_INFO, SEND_CODE_SUCCESS, str)
	} else {
		SendConnMessageJson(conn, PMD_ROOM_INFO, SEND_CODE_ERROR, "不存在此房间")
	}
}

func NewRoomLogic() Roomer {
	return &room{}
}
