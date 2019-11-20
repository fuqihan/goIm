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

	PMD_ROOM_CREATE int = 3001

	PMD_ROOM_JOIN int = 3002

	PMD_ROOM_QUIT int = 3003

	PMD_ROOM_INFO int = 3004
)

const (
	// 请求返回码
	SEND_CODE_SUCCESS int32 = 1

	SEND_CODE_ERROR int32 = 2
)

const (
	//redis key
	REDIS_ROOM_MESSAGE string = "room:message:%d"

	REDIS_ROOM_LIST string = "room:list" // set

	REDIS_ROOM_NAME_LIST string = "room:name:list"

	REDIS_ROOM_DETAIL string = "room:detail:%d" // hash

	REDIS_ROOM_ROLE string = "room:role:%d:%d"

	REDIS_ROOM_USERS string = "room:users:%d"

	REDIS_ROOM_USER_INFO string = "room:user:%d:%s"

	REDIS_USER_ROOMS string = "user:rooms:%s"

	REDIS_USER_SWAP_LIST string = "user:swap:list:%s" // 每个人交流的列表

	REDIS_USER_SWAP_DETAIL string = "user:swap:%S:%S"

	REDIS_USER_SINGLE_SEND string = "user:single:send:%s:%s" // 两个人之间的交流记录
)

const (
	// error 内容
	ERROR_TEXT_PARAM string = "参数不正确或为空"
)

const (
	// 房间类型
	ROOM_MODE_DEFAULT uint32 = 1 // 普通类型
)

const (
	// 房间权限
	ROOM_ROLE_SADMIN uint32 = 1001 // 系统超管
	ROOM_ROLE_ADMIN  uint32 = 1    // 群超管
)

const (
	// 用户间关系
	USER_STATUS_STRANGE uint32 = 1 // 陌生人
	USER_STATUS_FRIEND  uint32 = 2 // 好友
	USER_STATUS_BLOCK   uint32 = 3 // 拉黑
)
