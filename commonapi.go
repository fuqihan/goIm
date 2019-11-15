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

/*
	群聊发送
*/
type SendRoomMessageApi struct {
	UserId string `json:"user_id"`
	RoomId string `json:"room_id"`
	Str    string `json:"str"`
	Now    uint   `json:"now"`
}

type CreateRoomApi struct {
	UserId   string `json:"user_id"`
	RoomName string `json:"room_name"`
}

/*
	加入房间
*/
type JoinRoomApi struct {
	UserIds []string `json:"user_ids"` // 支持十足模式
	RoomId  string   `json:"room_id"`  // 房间id
}

/*
	退出房间
*/
type QuitRoomApi struct {
	UserIds []string `json:"user_ids"`
	RoomId  string   `json:"room_id"`
}

type GetRoomInfoApi struct {
	RoomId string `json:"room_id"`
}

/*
	回执请求
*/
type SendReceiptApi struct {
	Mode uint32 `json:"mode"` // 模式 2 为单聊 3 为群聊  4 为系统通知
	Form string `json:"form"` // 回执的发送者
	To   string `json:"to"`   // 回执的接收者
	Now  uint   `json:"now"`  // 当前看到的时间
}
