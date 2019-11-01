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

	// Error 枚举
	ERROR_PARSE_JSON int = 1001 // json解析错误
)
