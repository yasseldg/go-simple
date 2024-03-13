package sJson

import (
	"encoding/json"

	"github.com/yasseldg/go-simple/logs/sLog"
)

func ToObj(msg string, obj any) error {
	err := json.Unmarshal([]byte(msg), obj)
	if err != nil {
		sLog.Error("json.Unmarshal([]byte(msg), obj): %s", err)
		return err
	}
	return nil
}

func ToString(v interface{}) (string, error) {
	result, err := json.Marshal(v)
	if err != nil {
		sLog.Error("json.Marshal(v): %s", err)
		return "", err
	}
	return string(result), nil
}
