package internal

import "strings"

type Environment string

const (
	// Environment
	Env_Write = Environment("write")
	Env_Read  = Environment("read")
)

func GetEnvironment(env string) Environment {
	switch strings.ToLower(env) {
	case "write":
		return Env_Write

	default:
		return Env_Read
	}
}
