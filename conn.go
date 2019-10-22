package goIm

import (
	"fmt"
	"log"
	"net"
	"strings"
)

/**
接受信息处理
*/
func read(c net.Conn) {
	// TODO 添加字节长度配置
	data := make([]byte, 1000)
	for {
		n, err := c.Read(data)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(data[0:n]))
		go ConnHandle(c, string(data[0:n]))
	}
}

/**
发送信息
*/
func SendConnMessage(c net.Conn, str string) {
	if _, err := c.Write([]byte(str)); err != nil {
		fmt.Println(err)
	}
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
		go read(c)
	}
}
