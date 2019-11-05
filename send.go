package goIm

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type SendApi struct {
	Pmd  int    `json:"pmd"`
	Data string `json:"data"`
	Code int32  `json:"code"`
}

/*
	发送信息转成json
*/
func SendConnMessageJson(c net.Conn, pmd int, code int32, str string) error {
	m := SendApi{Pmd: pmd, Data: str, Code: code}
	if jsonBytes, err := json.Marshal(m); err == nil {
		SendConnMessage(c, jsonBytes)
	} else {
		fmt.Println(err)
	}
	return nil
}

func SendConnMessageInt(c net.Conn, i interface{}) {

}

/*
	信息参数为string，节省转换消耗
*/
func SendConnMessageStr(c net.Conn, str string) {
	SendConnMessage(c, []byte(str))
}

/**
发送信息,
*/
func SendConnMessage(c net.Conn, b []byte) error {
	length := len(b)
	lenNum := make([]byte, 2)
	binary.BigEndian.PutUint16(lenNum, uint16(length))
	pkg := new(bytes.Buffer)
	//写入长度
	if err := binary.Write(pkg, binary.BigEndian, lenNum); err != nil {
		return err
	}
	//写入消息体
	if err := binary.Write(pkg, binary.BigEndian, b); err != nil {
		return err
	}

	if _, err := c.Write(pkg.Bytes()); err != nil {
		return err
	}
	return nil
}
