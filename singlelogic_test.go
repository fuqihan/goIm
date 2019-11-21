package goIm

import (
	"fmt"
	"goIm/utils"
	"net"
	"testing"
	"time"
)

func TestSingle_SendMessage(t *testing.T) {
	if conn, err := InitTestDial("3000"); err == nil {
		go ReadConnMessage(conn, func(conn net.Conn, s string) {
			obj := new(SendApi)
			utils.ParseJson(s, obj)
			if obj.Code != SEND_CODE_SUCCESS {
				t.Errorf("返回错误 %s, 应该返回", s)
			}
			return
		})
		obj := make(map[string]interface{})
		obj["to"] = "111111"
		obj["form"] = "22222"
		obj["str"] = "我们是冠军"
		obj["now"] = utils.GetTimeNow()
		joinStr := fmt.Sprintf(`{"pmd": %d, "data": {"to": "%s", "form": "%s", "str": "%s", "now": %d}}`,
			PMD_SINGLE_SEND_MESSAGE, obj["to"], obj["form"], obj["str"], obj["now"])
		SendConnMessageStr(conn, joinStr)
		time.Sleep(5e9)
	}
}

func TestSingle_SendReceipt(t *testing.T) {
	if conn, err := InitTestDial("3000"); err == nil {
		go ReadConnMessage(conn, func(conn net.Conn, s string) {
			obj := new(SendApi)
			utils.ParseJson(s, obj)
			if obj.Code != SEND_CODE_SUCCESS {
				t.Errorf("返回错误 %s, 应该返回", s)
			}
			return
		})
		obj := make(map[string]interface{})
		obj["to"] = "111111"
		obj["form"] = "22222"
		obj["now"] = utils.GetTimeNow()
		joinStr := fmt.Sprintf(`{"pmd": %d, "data": {"to": "%s", "form": "%s", "now": %d}}`,
			PMD_SINGLE_RECEIPT, obj["to"], obj["form"], obj["now"])
		SendConnMessageStr(conn, joinStr)
		time.Sleep(5e9)
	}
}
