/* 全局对象 */
package goIm

import (
	"goIm/dbredis"
	"goIm/utils"
)

var (
	DBRedisConn dbredis.RedisHandler // redis 数据库操作器
	RoomUUIDGen *utils.UUIDGenerator
)
