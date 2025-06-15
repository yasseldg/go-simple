package tInterval

import (
	"fmt"
	"time"
)

type (
	Interval string

	Intervals []Interval
)

const (
	Interval_1m  = Interval("1m")
	Interval_5m  = Interval("5m")
	Interval_15m = Interval("15m")
	Interval_30m = Interval("30m")
	Interval_1h  = Interval("1h")
	Interval_4h  = Interval("4h")
	Interval_D   = Interval("D")
	Interval_W   = Interval("W")

	// long intervals
	Interval_M = Interval("M")
	Interval_Y = Interval("Y")

	DEFAULT = Interval("DEFAULT")
)

func (i Interval) IsDefault() bool {
	return i == DEFAULT
}

func (i Interval) String() string {
	if i.IsDefault() {
		return ""
	}
	return string(i)
}

func (i Interval) Get() Interval {
	return i
}

func (i Interval) ForLog() string {
	return fmt.Sprintf(" ( %s ) ", i.String())
}

func (i Interval) Minutes() int64 {
	switch i {
	case Interval_1m:
		return 1
	case Interval_5m:
		return 5
	case Interval_15m:
		return 15
	case Interval_30m:
		return 30
	case Interval_1h:
		return 60
	case Interval_4h:
		return 240
	case Interval_D:
		return 1440
	case Interval_W:
		return 10080
	default:
		return 0
	}
}

func (i Interval) Seconds() int64 {
	return i.Minutes() * 60
}

func (i Interval) MilliSeconds() int64 {
	return i.Seconds() * 1000
}

// IsClosing returns true if the timestamp in inter is the closing time of the interval
func (i Interval) IsClosing(ts int64, inter Inter) bool {
	if i.IsDefault() || inter.IsDefault() {
		return false
	}
	return i.Next(ts) == inter.Next(ts)
}

func (i Interval) Prev(ts int64) int64 {
	if ts <= 0 {
		return 0
	}

	switch i {
	case DEFAULT:
		return 0

	case Interval_M:
		return startOfMonth(ts)

	case Interval_Y:
		return startOfYear(ts)

	default:
		return ts - (ts % i.Seconds())
	}
}

func (i Interval) Next(ts int64) int64 {
	if ts <= 0 {
		return 0
	}

	switch i {
	case DEFAULT:
		return 0

	case Interval_M:
		return startOfNextMonth(ts)

	case Interval_Y:
		return startOfNextYear(ts)

	default:
		return i.Prev(ts) + i.Seconds()
	}
}

func (i Interval) AsaName(ts int64) string {
	switch i {
	case Interval_M:
		return time.Unix(ts, 0).Format("Jan 06")

	case Interval_Y:
		return time.Unix(ts, 0).Format("06")

	default:
		return ""
	}
}

func (i Interval) Duration() time.Duration {
	return time.Duration(i.Seconds()) * time.Second
}
