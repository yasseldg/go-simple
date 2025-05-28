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
	IsActive() bool
	Active() error
	Inactive()

	SetLocation(string) error
	SetMinute(string) error
	SetHour(string) error
	SetDay(string) error
	SetMonth(string) error
	SetDayOfWeek(string) error
	SetSchedule(string) error

	Next(int64) int64
}
