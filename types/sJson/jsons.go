package sJson

import (
	"encoding/json"
)

func ToObj(msg string, obj any) error {
	return ByteToObj([]byte(msg), obj)
}

func ByteToObj(data []byte, obj any) error {
	err := json.Unmarshal(data, obj)
	if err != nil {
		return err
	}
	return nil
}

func ToString(v interface{}) (string, error) {
	result, err := ToByte(v)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func ToByte(v interface{}) ([]byte, error) {
	result, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return result, nil
}
