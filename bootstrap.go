package goIm

import (
	"fmt"
	"goIm/dbredis"
	"log"
	"net"
	"strings"
)

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
	dbRedisBootstrap()
	for {
		c, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go ReadConnMessage(c, ConnHandle)
	}
}

func dbRedisBootstrap() {
	DbRedisConn, _ = dbredis.CreateConn()
}
