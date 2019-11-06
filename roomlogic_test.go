package goIm

import (
	"fmt"
	"goIm/utils"
	"net"
	"testing"
	"time"
)

func TestRoom_Join(t *testing.T) {
	if conn, err := InitTestDial("3000"); err == nil {
		userId := "aaaa"
		roomName := "asasass"
		go ReadConnMessage(conn, func(conn net.Conn, s string) {
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
