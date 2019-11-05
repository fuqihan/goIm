package utils

import "time"

var (
	//东八区
	_CST_ZONE = time.FixedZone("CST", 8*3600)
)

func GetTimeNow() int64 {
	return time.Now().In(_CST_ZONE).Unix()

}
