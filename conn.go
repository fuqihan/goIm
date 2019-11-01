package goIm

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strings"
)

var (
	_lastReadByfes = make([]byte, 0) // 上次解包剩余的数据
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
		if length := len(_lastReadByfes); length != 0 {
			result.Write(_lastReadByfes)
			_lastReadByfes = make([]byte, 0)
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
			_lastReadByfes = data[0:]
		}
	}
	return
}

/**
发送信息
*/
func SendConnMessage(c net.Conn, str string) error {
	msg := []byte(str)
	length := len(msg)
	lenNum := make([]byte, 2)
	binary.BigEndian.PutUint16(lenNum, uint16(length))
	pkg := new(bytes.Buffer)
	//写入长度
	if err := binary.Write(pkg, binary.BigEndian, lenNum); err != nil {
		return err
	}
	//写入消息体
	if err := binary.Write(pkg, binary.BigEndian, msg); err != nil {
		return err
	}

	if _, err := c.Write(pkg.Bytes()); err != nil {
		return err
	}
	return nil
}

/**
启动类
*/
func Bootstrap(port string) error {
	//strings.Join()
	urls := []string{"127.0.0.1:", port}
	log.Println(port, "tcp启动")
	// TODO 换成ListenTCP 开启keep-alive
	listen, err := net.Listen("tcp", strings.Join(urls, ""))
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Println(port, "tcp启动成功")
	for {
		c, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go ReadConnMessage(c, ConnHandle)
	}
}
