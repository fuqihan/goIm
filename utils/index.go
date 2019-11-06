package utils

import "time"

var (
	//东八区
	_CST_ZONE = time.FixedZone("CST", 8*3600)
)

/*
	[key, val, key, val]
*/
type MaptoArr struct {
	Arr []interface{}
}

func (ma *MaptoArr) Add(args ...interface{}) {
	ma.Arr = append(ma.Arr, args...)
}

func GetTimeNow() int64 {
	return time.Now().In(_CST_ZONE).Unix()

}
