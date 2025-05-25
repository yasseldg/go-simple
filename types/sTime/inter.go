package sTime

import "time"

type InterProvider interface {
	SetComma(rune)
	SetLayout(string)
	SetLayoutDate(string)
	SetLayoutTime(string)
	SetLocation(string) error
	SetTimeDiff(time.Duration)

	GetUTC(string) (int64, error)

	Comma() rune
	LayoutDate() string
	LayoutTime() string
}

type InterCron interface {
	Next(int64) int64
}
