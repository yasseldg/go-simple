package sTime

import "time"

type InterProvider interface {
	SetComma(rune)
	SetLayout(string)
	SetLayoutDate(string)
	SetLayoutTime(string)
	SetLocation(string) error
	SetTimeDiff(time.Duration)

	Comma() rune
	LayoutDate() string
	LayoutTime() string

	GetUTC(string) (int64, error)
	GetDate(int64) string
	GetTime(int64) string
	GetDateAndTime(int64) (string, string)
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
