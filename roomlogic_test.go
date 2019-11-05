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
				t.Errorf("返回错误 %s, 应该返回 %d", s, ERROR_PARSE_JSON)
			}
			return
		})
		IniTestLogin(conn, userId)
		joinStrs := fmt.Sprintf(`{"pmd": %d, "data": {"userId": "%s", "roomName": "%s", "type": 1}}`,
			PMD_ROOM_JOIN, userId, roomName)
		SendConnMessageStr(conn, joinStrs)
		time.Sleep(1e9)
	}
}
