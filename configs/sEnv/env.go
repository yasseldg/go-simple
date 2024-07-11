package sEnv

import (
	"os"
	"strconv"

	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sInts"
	"github.com/yasseldg/go-simple/types/sStrings"

	"gopkg.in/yaml.v3"
)

// Get
func Get(env_name, defaults string) string {
	env, ok := os.LookupEnv(env_name)
	if ok {
		return env
	}
	return defaults
}

// GetSlice
func GetSlice(env_name string, defaults ...string) []string {
	string_values := Get(env_name, "")
	if len(string_values) > 0 {
		return sStrings.SplitString(string_values, ",")
	}
	return defaults
}

// GetInt
func GetInt(env_name string, defaults int) int {
	env, ok := os.LookupEnv(env_name)
	if ok {
		return sInts.Get(env)
	}
	return defaults
}

// GetBool
func GetBool(env_name string, defaults bool) bool {
	str := Get(env_name, "")
	b, err := strconv.ParseBool(str)
	if err != nil {
		return defaults
	}
	return b
}

// GetSliceInt
func GetSliceInt(env_name string, defaults ...int) (res []int) {
	values := GetSlice(env_name)
	if len(values) > 0 {
		for _, v := range values {
			res = append(res, sInts.Get(v))
		}
		return
	}
	return defaults
}

// GetFloat64
func GetFloat64(env_name string, defaults float64) float64 {
	env, ok := os.LookupEnv(env_name)
	if ok {
		return sFloats.Get64(env)
	}
	return defaults
}

// LoadYaml
func LoadYaml(file_path string, obj interface{}) error {
	data, err := os.ReadFile(file_path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, obj)
	if err != nil {
		return err
	}
	return nil
}
