package goIm

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type Singler interface {
	SendMessage(conn net.Conn, userId string, message *SendMessageApi) error
	SendReceipt(conn net.Conn, msg *SendReceiptApi) error
}

type single struct {
}

func NewSingleLogic() Singler {
	return &single{}
}

func (s *single) SendMessage(conn net.Conn, userId string, message *SendMessageApi) error {

	if message.To != "" && message.Str != "" {
		message.Form = userId
		message.Now = int32(time.Now().Unix())
		if str, err := json.Marshal(message); err == nil {
			SendUserMessage(message.To, string(str))
		} else {
			fmt.Println(err)
		}
	}
	return nil
}

func (s *single) SendReceipt(conn net.Conn, msg *SendReceiptApi) error {

	return nil
}
