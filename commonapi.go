package goIm

type SendMessageApi struct {
	To   string `json:"-"`    // 接收者
	Form string `json:"form"` // 发送者
	Str  string `json:"str"`  // 发送的文件
	now  int32  `json:"now"`  // 时间戳
}

type JoinRoomApi struct {
	UserId   string   `json:"user_id"`
	UserIds  []string `json:"user_ids"`
	RoomName string   `json:"room_name"`
}
