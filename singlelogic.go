package goIm

import (
	"encoding/json"
	"fmt"
	"net"
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

	if message.To != "" && message.Str != "" && message.Now != 0 {
		message.Form = userId
		if str, err := json.Marshal(message); err == nil {
			SendUserMessage(message.To, string(str))
		} else {
			fmt.Println(err)
		}
	}
	SendConnMessageJson(conn, PMD_SINGLE_SEND_MESSAGE, SEND_CODE_ERROR, ERROR_TEXT_PARAM)
	return nil
}

func (s *single) SendReceipt(conn net.Conn, msg *SendReceiptApi) error {

	return nil
}
