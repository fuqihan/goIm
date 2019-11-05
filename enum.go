package goIm

const (
	// Pmd 枚举
	// 1 开头为用户账号相关
	// 2 开头为单聊操作
	// 3 开头为群聊操作
	// 4 开头为管理员操作操作
	PMD_LOGIN int = 1001

	PMD_SINGLE_SEND_MESSAGE int = 2001

	PMD_SINGLE_RECEIPT int = 2002

	PMD_ROOM_JOIN int = 3001

	// 请求返回码
	SEND_CODE_SUCCESS int32 = 1

	SEND_CODE_ERROR int32 = 2

	// Error 枚举
	ERROR_PARSE_JSON int = 1001 // json解析错误

	// error 内容
	ERROR_TEXT_PARAM string = "参数不正确或为空"

	// 房间类型
	ROOM_MODE_DEFAULT uint32 = 1 // 普通类型

	// 房间权限
	ROOM_ROLE_SADMIN uint32 = 1001 // 系统超管
	ROOM_ROLE_ADMIN  uint32 = 1    // 群超管
)
