package sStrings

import "fmt"

type Color int

const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
)

func Colored(color Color, msg string) string {
	return fmt.Sprintf("\033[%dm%s\033[0m", color, msg)
}
