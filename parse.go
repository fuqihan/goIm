package goIm

import (
	"encoding/json"
)

func ParseJson(str string, m interface{}) error {
	if err := json.Unmarshal([]byte(str), m); err != nil {
		return err
	}
	return nil
}
