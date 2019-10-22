package goIm

import (
	"encoding/json"
	"fmt"
)

func ParseJson(str string, m interface{}) {
	if err := json.Unmarshal([]byte(str), m); err != nil {
		fmt.Println(err)
	}
}
