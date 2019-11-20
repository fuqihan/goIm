package utils

import "time"

var (
	//东八区
	_CST_ZONE = time.FixedZone("CST", 8*3600)
)

/*
	[key, val, key, val]
*/
type maptoArr struct {
	Arr []interface{}
}

func (ma *maptoArr) Add(args ...interface{}) {
	ma.Arr = append(ma.Arr, args...)
}

func (ma *maptoArr) Clone() {
	ma.Arr = []interface{}{}
}

func NewMapToArr() *maptoArr {
	return &maptoArr{Arr: []interface{}{}}
}

func GetTimeNow() int64 {
	return time.Now().In(_CST_ZONE).Unix()

}
