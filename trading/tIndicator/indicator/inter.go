package indicator

import "github.com/yasseldg/go-simple/data/dIter"

type Inter interface {
	Count() int
	Increase()
}

type InterPeriodsConfig interface {
	dIter.Inter

	Periods() int
	Count() int
	Reset()
	Next() bool
}
