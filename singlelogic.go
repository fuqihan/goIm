package goIm

import (
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
	"goIm/utils"
	"net"
	"sort"
)

type Singler interface {
	SendMessage(conn net.Conn, data interface{}) // 单聊
	SendReceipt(conn net.Conn, data interface{}) // 消息回执
	//ChangeUserStatus(conn net.Conn, data interface{}) // 拉黑/取消拉黑

}

type single struct {
}

func NewSingleLogic() Singler {
	return &single{}
}

func (s *single) SendMessage(conn net.Conn, data interface{}) {
	message := new(SendMessageApi)
	mapstructure.Decode(data, message)
	if message.To != "" && message.Str != "" && message.Form != "" && message.Now != 0 {
		if str, err := json.Marshal(message); err == nil {
			ids := []string{message.To, message.Form}
			sort.Strings(ids)
			DBRedisConn.DoSetArgs("ZADD",
				fmt.Sprintf(REDIS_USER_SINGLE_SEND, ids[0], ids[1]), message.Now, string(str))
			//	初始化操作
			if DBRedisConn.DoSismember(fmt.Sprintf(REDIS_USER_SWAP_LIST, ids[0]), ids[1]) == 0 {
				DBRedisConn.DoSet("SADD",
					fmt.Sprintf(REDIS_USER_SWAP_LIST, ids[0]), ids[1])
				DBRedisConn.DoSet("SADD",
					fmt.Sprintf(REDIS_USER_SWAP_LIST, ids[1]), ids[0])
				now := utils.GetTimeNow()
				ma := utils.NewMapToArr()
				ma.Add("createDate", now)
				ma.Add("status", USER_STATUS_STRANGE)
				DBRedisConn.DoSetArgs("HMSET",
					fmt.Sprintf(REDIS_USER_SWAP_DETAIL, ids[0], ids[1]), ma.Arr...)

			}
			SendUserMessage(message.To, string(str))
		} else {
			fmt.Println(err)
		}
	}
	SendConnMessageJson(conn, PMD_SINGLE_SEND_MESSAGE, SEND_CODE_ERROR, ERROR_TEXT_PARAM)
}

func (s *single) SendReceipt(conn net.Conn, data interface{}) {
	msg := new(SendReceiptApi)
	mapstructure.Decode(data, msg)
	ids := []string{msg.To, msg.Form}
	sort.Strings(ids)
	var arrs []string
	if ids[0] == msg.Form {
		arrs = []string{"firstCurrentReceipt", string(msg.Now)}
	} else {
		arrs = []string{"twoCurrentReceipt", string(msg.Now)}
	}
	DBRedisConn.DoSetArgs("HSET",
		fmt.Sprintf(REDIS_USER_SWAP_DETAIL, ids[0], ids[1]), arrs[0], arrs[1])
}
