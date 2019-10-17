package goIm

import (
	"encoding/json"
	"log"
)

func ParseJson(str string, m interface{}) {
	if err := json.Unmarshal([]byte(str), m); err != nil {
		log.Fatal(err)
	}
}