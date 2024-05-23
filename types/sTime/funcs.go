package sTime

import (
	"fmt"
	"time"

	"github.com/yasseldg/go-simple/logs/sLog"
)

func Since(duration time.Duration) string {
	if duration < 0 {
		return "" // o manejar de otra manera si se requiere
	}

	if duration == 0 {
		return "0s"
	}

	var s string
	hours := duration / time.Hour
	if hours > 0 {
		s += fmt.Sprintf("%dh ", hours)
		duration -= hours * time.Hour
	}

	minutes := duration / time.Minute
	if minutes > 0 || hours > 0 {
		s += fmt.Sprintf("%dm ", minutes)
		duration -= minutes * time.Minute
	}

	seconds := duration / time.Second
	if seconds > 0 || minutes > 0 || hours > 0 {
		s += fmt.Sprintf("%ds ", seconds)
	}

	return s
}

func TimeControl(f func(), name string) {
	t := time.Now()

	msg := fmt.Sprintf("Time Control ( %s )", name)

	sLog.Info(msg)

	f()

	fmt.Println()

	sLog.Info(fmt.Sprintf("%s elapsed %s \n", msg, Since(time.Since(t))))
}
