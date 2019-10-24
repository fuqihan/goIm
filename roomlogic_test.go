package goIm

import (
	"fmt"
	"testing"
	"time"
)

func TestRoom_Join(t *testing.T) {

	if conn, err := InitTestDial("3000"); err == nil {
		userId := "aaaa"
		roomName := "asasass"
		IniTestLogin(conn, userId)
		joinStrs := fmt.Sprintf(`{"pmd": %d, "data": {"userId": "%s", "roomName": "%s"}}`,
			PMD_ROOM_JOIN, userId, roomName)
		SendConnMessage(conn, joinStrs)
		time.Sleep(1e9)
	}
}
