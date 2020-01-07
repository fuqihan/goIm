package goIm

import (
	"fmt"
	"goIm/dbredis"
	"goIm/utils"
	"log"
	"net"
)

const (
	_defaultPort uint = 3000
)

const (
	PARSE_IOTA = iota
	PARSE_JSON
	PARSE_PROTO
)

type IMOptions struct {
	redisConfig *dbredis.ConnOptions
	port        uint
	address     string
	ssl         bool
	parseClass  uint // json or proto default json
}

func NewIMOptions() *IMOptions {
	c := new(IMOptions)
	c.setPtrs()
	return c
}

func (op *IMOptions) setPtrs() {
	op.redisConfig = new(dbredis.ConnOptions)
	op.port = _defaultPort
	op.address = "0.0.0.0"
	op.ssl = false
	op.parseClass = PARSE_JSON
}

/**
启动类
*/
func Bootstrap(op *IMOptions) error {
	//strings.Join()
	address := fmt.Sprintf("%s:%d", op.address, op.port)
	log.Println(op.port, "tcp启动")
	// TODO 换成ListenTCP 开启keep-alive
	listen, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.Println(op.port, "tcp启动成功")
	dbRedisBootstrap()
	UUIDBootstrap()
	for {
		c, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go ReadConnMessage(c, op, ConnHandle)
	}
}

func dbRedisBootstrap() {
	redisOp := new(dbredis.ConnOptions)
	redisOp.Host = "49.235.242.138"
	redisOp.Port = 6379
	redisOp.Password = "aliyunsb"
	DBRedisConn, _ = dbredis.CreateConn(redisOp)
}

func UUIDBootstrap() {
	RoomUUIDGen = utils.NewUUIDGenerator()
}
