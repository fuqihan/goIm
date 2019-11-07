package goIm

/*
	发送信息
*/
type SendMessageApi struct {
	To   string `json:"-"`    // 接收者
	Form string `json:"form"` // 发送者
	Str  string `json:"str"`  // 发送的文件
	Now  uint   `json:"now"`  // 时间戳
}

type CreateRoomApi struct {
	UserId   string `json:"user_id"`
	RoomName string `json:"room_name"`
}

/*
	加入房间
*/
type JoinRoomApi struct {
	UserId   string   `json:"user_id"`   // 要加入房间的id
	UserIds  []string `json:"user_ids"`  // 支持十足模式
	RoomName string   `json:"room_name"` // 房间名
}

/*
	退出房间
*/
type QuitRoomApi struct {
	UserId   string `json:"user_id"`
	RoomName string `json:"room_name"`
}

/*
	回执请求
*/
type SendReceiptApi struct {
	Mode uint32 `json:"mode"` // 模式 2 为单聊 3 为群聊  4 为系统通知
	Form string `json:"form"` // 回执的发送者
	Now  uint   `json:"now"`  // 当前看到的时间
}
