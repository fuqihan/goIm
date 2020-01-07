package goIm

import (
	"fmt"
	"goIm/utils"
	"net"
	"testing"
	"time"
)

func TestRoom_Create(t *testing.T) {
	op := NewIMOptions()
	if conn, err := InitTestDial("3000", op); err == nil {
		userId := "aaaa"
		roomName := "asasass"
		go ReadConnMessage(conn, op, func(conn net.Conn, s string, op *IMOptions) {
			obj := new(SendApi)
			utils.ParseJson(s, obj)
			if obj.Code != SEND_CODE_SUCCESS {
				t.Errorf("返回错误 %s, 应该返回", s)
			}
			return
		})
		IniTestLogin(conn, userId)
		joinStr := fmt.Sprintf(`{"pmd": %d, "data": {"userId": "%s", "roomName": "%s"}}`,
			PMD_ROOM_CREATE, userId, roomName)
		SendConnMessageStr(conn, joinStr)
		time.Sleep(1e9)
	}
}

func TestRoom_Join(t *testing.T) {
	op := NewIMOptions()
	if conn, err := InitTestDial("3000", op); err == nil {
		userId := "aaaa"
		go ReadConnMessage(conn, op, func(conn net.Conn, s string, op *IMOptions) {
			obj := new(SendApi)
			utils.ParseJson(s, obj)
			if obj.Code != SEND_CODE_SUCCESS {
				t.Errorf("返回错误 %s, 应该返回", s)
			}
			return
		})
		IniTestLogin(conn, userId)
		joinStr := fmt.Sprintf(`{"pmd": %d, "data": {"roomId": "%d", "userIds": ["%s"]}}`,
			PMD_ROOM_JOIN, 1, "roomName")
		SendConnMessageStr(conn, joinStr)
		time.Sleep(1e9)
	}
}

func TestRoom_Quit(t *testing.T) {
	op := NewIMOptions()
	if conn, err := InitTestDial("3000", op); err == nil {
		userId := "aaaa"
		go ReadConnMessage(conn, op, func(conn net.Conn, s string, op *IMOptions) {
			obj := new(SendApi)
			utils.ParseJson(s, obj)
			if obj.Code != SEND_CODE_SUCCESS {
				t.Errorf("返回错误 %s, 应该返回", s)
			}
			return
		})
		//IniTestLogin(conn, userId)
		joinStr := fmt.Sprintf(`{"pmd": %d, "data": {"roomId": "%d", "userIds": ["%s"]}}`,
			PMD_ROOM_QUIT, 1, userId)
		SendConnMessageStr(conn, joinStr)
		time.Sleep(1e9)
	}
}

func TestRoom_GetRoomInfo(t *testing.T) {
	op := NewIMOptions()
	if conn, err := InitTestDial("3000", op); err == nil {
		go ReadConnMessage(conn, op, func(conn net.Conn, s string, op *IMOptions) {
			obj := new(SendApi)
			utils.ParseJson(s, obj)
			fmt.Println(obj)
			if obj.Code != SEND_CODE_SUCCESS {
				t.Errorf("返回错误 %s, 应该返回", s)
			}
			return
		})
		joinStr := fmt.Sprintf(`{"pmd": %d, "data": {"roomId": "%d"}}`,
			PMD_ROOM_INFO, 1)
		SendConnMessageStr(conn, joinStr)
		time.Sleep(1e9)
	}
}

func TestRoom_SendMessage(t *testing.T) {
	op := NewIMOptions()
	if conn, err := InitTestDial("3000", op); err == nil {
		go ReadConnMessage(conn, op, func(conn net.Conn, s string, op *IMOptions) {
			obj := new(SendApi)
			utils.ParseJson(s, obj)
			fmt.Println(obj)
			if obj.Code != SEND_CODE_SUCCESS {
				t.Errorf("返回错误 %s, 应该返回", s)
			}
			return
		})
		userId := "aaaa1"
		now := utils.GetTimeNow()
		joinStr := fmt.Sprintf(`{"pmd": %d, "data": {"roomId": %d, "userId", "%s", str: "asasasas", now: %d}}`,
			PMD_ROOM_SEND_MESSAGE, 1, userId, now)
		SendConnMessageStr(conn, joinStr)
		time.Sleep(1e9)
	}
}
