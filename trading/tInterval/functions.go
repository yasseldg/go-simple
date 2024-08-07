package tInterval

import (
	"strings"
	"time"

	"github.com/yasseldg/go-simple/types/sInts"
)

func Get(interval string) Interval {
	switch interval {
	case "M1", "candle1m", "1m", "1", "1min":
		return Interval_1m
	case "M5", "candle5m", "5m", "5", "5min":
		return Interval_5m
	case "M15", "candle15m", "15m", "15", "15min":
		return Interval_15m
	case "M30", "candle30m", "30m", "30", "30min":
		return Interval_30m
	case "H1", "candle1H", "1H", "60", "1h":
		return Interval_1h
	case "H4", "candle4H", "4H", "240", "4h":
		return Interval_4h
	case "D", "candle1D", "1D", "1Dutc", "1440", "1d":
		return Interval_D
	case "W", "candle1W", "1W", "Week", "10080", "1w":
		return Interval_W
	case "M", "Month":
		return Interval_M
	default:
		return Interval_DEFAULT
	}
}

func GetMult(interval string) (Interval, int) {
	mult := 1
	strs := strings.Split(interval, "*")
	if len(strs) > 1 {
		interval = strs[0]
		mult_2 := sInts.Get(strs[1])
		if mult_2 > mult {
			mult = mult_2
		}
	}
	return Get(interval), mult
}

func IsSameMonth(ts1, ts2 int64) bool {
	t1 := time.Unix(ts1, 0).UTC()
	t2 := time.Unix(ts2, 0).UTC()
	return t1.Year() == t2.Year() && t1.Month() == t2.Month()
}

func startOfMonth(ts int64) int64 {
	t := time.Unix(ts, 0).UTC()
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC).Unix()
}

func startOfNextMonth(ts int64) int64 {
	t := time.Unix(ts, 0).UTC()
	y, m, _ := t.Date()
	if m == 12 {
		y++
		m = 0
	}
	return time.Date(y, m+1, 1, 0, 0, 0, 0, time.UTC).Unix()
}

func GetIntervals(intervals ...string) InterIterLimited {
	iter := NewIterLimited()

	for _, interval := range intervals {
		_interval := Get(interval)
		if _interval.IsDefault() {
			continue
		}

		iter.Add(_interval)
	}

	return iter
}
