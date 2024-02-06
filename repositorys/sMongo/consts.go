package sMongo

type Environment string

const (
	// Environment
	EnvironmentWrite = Environment("write")
	EnvironmentRead  = Environment("read")
)

func getEnvironment(env string) Environment {
	switch env {
	case "write":
		return EnvironmentWrite

	default:
		return EnvironmentRead
	}
}
