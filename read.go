package goIm

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

var (
	_lastReadBytes = make([]byte, 0) // 上次解包剩余的数据
)

/**
接受信息处理
*/
func ReadConnMessage(c net.Conn, fn func(conn net.Conn, s string)) {
	// TODO 添加字节长度配置
	data := make([]byte, 1000)
	result := bytes.NewBuffer(nil)
	for {
		n, err := c.Read(data)
		if err != nil {
			fmt.Println(err)
			break
		}
		if length := len(_lastReadBytes); length != 0 {
			result.Write(_lastReadBytes)
			_lastReadBytes = make([]byte, 0)
		}
		result.Write(data[0:n])
		scanner := bufio.NewScanner(result)
		scanner.Split(packetSplitFunc)
		for scanner.Scan() {
			//fmt.Println(string(scanner.Bytes()))
			go fn(c, string(scanner.Bytes()))
		}
	}
}

func packetSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if !atEOF && len(data) > 2 {
		var length int16
		// 读出 数据包中 实际数据 的长度(大小为 0 ~ 2^16)
		binary.Read(bytes.NewReader(data[0:2]), binary.BigEndian, &length)
		pl := int(length) + 2
		if pl <= len(data) {
			return pl, data[2:pl], nil
		} else {
			_lastReadBytes = data[0:]
		}
	}
	return
}
